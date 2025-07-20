package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"teamfinder/backend/internal/services"
	"teamfinder/backend/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	emailService    *services.EmailService
	telegramService *services.TelegramService
	authService     *services.AuthService
	codes           map[string]codeInfo
}

type codeInfo struct {
	code      string
	createdAt time.Time
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type AuthResponse struct {
	AccessToken  string      `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string      `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User         UserProfile `json:"user"`
}

type UserProfile struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
}

type MessageResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"`
}

func NewAuthHandler(emailService *services.EmailService, telegramService *services.TelegramService) *AuthHandler {
	return &AuthHandler{
		emailService:    emailService,
		telegramService: telegramService,
		authService:     services.NewAuthService(),
		codes:           make(map[string]codeInfo),
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with username, email and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "User registration data"
// @Success      200 {object} AuthResponse "User registered successfully"
// @Failure      400 {object} ErrorResponse "Invalid request data"
// @Failure      409 {object} ErrorResponse "User already exists"
// @Failure      500 {object} ErrorResponse "Internal server error"
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "username") {
				c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			} else if strings.Contains(err.Error(), "email") {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			} else {
				c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			}
		} else {
			log.Printf("ERROR: Failed to create user '%s': %v", req.Username, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}

	// Создаем дефолтный профиль
	if err := h.authService.CreateDefaultProfile(user.ID, req.Username); err != nil {
		log.Printf("WARN: Failed to create default profile for user %d: %v", user.ID, err)
		// Не возвращаем ошибку, так как пользователь уже создан
	}

	accessToken, refreshToken, err := h.authService.GenerateTokens(user.ID, user.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for user %d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User registered successfully, id=%d, username='%s', email='%s'", user.ID, req.Username, req.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":       user.ID,
			"username": req.Username,
			"email":    req.Email,
		},
	})
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user with email and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "User login credentials"
// @Success      200 {object} AuthResponse "Login successful"
// @Failure      400 {object} ErrorResponse "Invalid request data"
// @Failure      401 {object} ErrorResponse "Invalid credentials"
// @Failure      500 {object} ErrorResponse "Internal server error"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.GetUserByEmail(req.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		log.Printf("ERROR: Database error during login for email '%s': %v", req.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	err = h.authService.ValidatePassword(user.PasswordHash, req.Password)
	if err != nil {
		log.Printf("WARN: Failed login attempt for email '%s'", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateTokens(user.ID, user.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for user %d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User logged in successfully, id=%d, email='%s'", user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *AuthHandler) SendEmailCode(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := h.emailService.GenerateVerificationCode()
	if err := h.emailService.SendVerificationCode(req.Email, code); err != nil {
		if strings.Contains(err.Error(), "please wait") {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Printf("ERROR: Failed to send verification code to '%s': %v", req.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send verification code",
		})
		return
	}

	h.codes[req.Email] = codeInfo{
		code:      code,
		createdAt: time.Now(),
	}

	log.Printf("INFO: Verification code sent to '%s'", req.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "Verification code sent successfully",
	})
}

func (h *AuthHandler) VerifyEmailCode(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Code     string `json:"code" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	codeInfo, exists := h.codes[req.Email]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No verification code found. Please request a new code."})
		return
	}

	if time.Since(codeInfo.createdAt) > time.Minute*15 {
		delete(h.codes, req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification code expired. Please request a new code."})
		return
	}

	if codeInfo.code != req.Code {
		log.Printf("WARN: Invalid verification code attempt for email '%s'", req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	user, err := h.authService.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		} else {
			log.Printf("ERROR: Failed to create user during email verification, username='%s': %v", req.Username, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}

	// Создаем дефолтный профиль
	if err := h.authService.CreateDefaultProfile(user.ID, req.Username); err != nil {
		log.Printf("WARN: Failed to create default profile for user %d: %v", user.ID, err)
	}

	accessToken, refreshToken, err := h.authService.GenerateTokens(user.ID, user.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens during email verification for user %d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	delete(h.codes, req.Email)

	log.Printf("INFO: User registered via email verification, id=%d, username='%s', email='%s'", user.ID, req.Username, req.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":       user.ID,
			"username": req.Username,
			"email":    req.Email,
		},
	})
}

func (h *AuthHandler) TelegramLogin(c *gin.Context) {
	authURL := h.telegramService.GenerateAuthURL()
	c.JSON(http.StatusOK, gin.H{
		"auth_url": authURL,
		"message":  "Open this URL in Telegram to authenticate",
	})
}

func (h *AuthHandler) VerifyTelegram(c *gin.Context) {
	var req map[string]string

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if !h.telegramService.ValidateAuthData(req) {
		log.Printf("WARN: Invalid Telegram auth data attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Telegram data"})
		return
	}

	telegramID := req["id"]
	username := req["username"]
	if username == "" {
		username = "telegram_user_" + telegramID
	}

	// TODO: Здесь можно создать пользователя или найти существующего по telegram_id
	// Пока используем моковый approach
	accessToken, refreshToken, err := utils.GenerateTokenPair(1, username+"@telegram.local")
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for Telegram auth: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User logged in via Telegram, telegram_id='%s', username='%s'", telegramID, username)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"telegram_id": telegramID,
			"username":    username,
		},
		"message": "Telegram authentication successful",
	})
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Get new access and refresh tokens using a valid refresh token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body RefreshTokenRequest true "Refresh token"
// @Success      200 {object} object{access_token=string,refresh_token=string} "Tokens refreshed successfully"
// @Failure      400 {object} ErrorResponse "Invalid request data"
// @Failure      401 {object} ErrorResponse "Invalid refresh token"
// @Failure      500 {object} ErrorResponse "Internal server error"
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидируем refresh token используя специальную функцию
	claims, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Проверяем существует ли пользователь
	user, err := h.authService.GetUserByID(int(claims.UserID))
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Генерируем новые токены
	accessToken, refreshToken, err := h.authService.GenerateTokens(int(claims.UserID), claims.Email)
	if err != nil {
		log.Printf("ERROR: Failed to refresh tokens for user %d: %v", claims.UserID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh tokens"})
		return
	}

	log.Printf("INFO: Tokens refreshed for user %d", claims.UserID)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// CheckToken godoc
// @Summary      Validate token
// @Description  Check if the provided JWT token is valid
// @Tags         Authentication
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} object{valid=bool,user_id=int,email=string,message=string} "Token is valid"
// @Failure      401 {object} ErrorResponse "Invalid or missing token"
// @Router       /auth/check [get]
func (h *AuthHandler) CheckToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"user_id": userID,
		"email":   email,
		"message": "Token is valid",
	})
}

// DeleteAccount godoc
// @Summary      Delete user account
// @Description  Permanently delete the user account and all associated data
// @Tags         Authentication
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} MessageResponse "Account deleted successfully"
// @Failure      401 {object} ErrorResponse "Invalid or missing token"
// @Failure      500 {object} ErrorResponse "Failed to delete account"
// @Router       /auth/account [delete]
func (h *AuthHandler) DeleteAccount(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDInt, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.authService.DeleteUser(int(userIDInt)); err != nil {
		log.Printf("ERROR: Failed to delete user %d: %v", userIDInt, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete account"})
		return
	}

	log.Printf("WARN: User account deleted, id=%d", userIDInt)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted successfully",
	})
}

// Logout godoc
// @Summary      User logout
// @Description  Logout user (client should remove tokens)
// @Tags         Authentication
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} MessageResponse "Logout successful"
// @Failure      401 {object} ErrorResponse "Invalid or missing token"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// В stateless JWT системе logout обычно происходит на клиенте
	// Можно добавить blacklist токенов в будущем
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful. Please remove tokens from your client.",
	})
}
