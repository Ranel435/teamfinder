package handlers

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strings"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
	"teamfinder/backend/internal/services"
	"teamfinder/backend/internal/utils"
	"time"

	"bytes"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	emailService    *services.EmailService
	telegramService *services.TelegramService
	codes           map[string]codeInfo
}
type codeInfo struct {
	code      string
	createdAt time.Time
}

// Add these new structs for request handling
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(emailService *services.EmailService, telegramService *services.TelegramService) *AuthHandler {
	return &AuthHandler{
		emailService:    emailService,
		telegramService: telegramService,
		codes:           make(map[string]codeInfo),
	}
}

// Вызывается внутри VerifyEmailCode, потом и в VerifyTelegram
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR: Password hashing failed for user '%s': %v", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	var userID int
	err = database.Pool.QueryRow(c, query, req.Username, req.Email, string(hashedPassword)).Scan(&userID)
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

	profileQuery := `
		INSERT INTO profiles (user_id, name, surname, hackathon_id)
		VALUES ($1, $2, $3, NULL)`

	_, err = database.Pool.Exec(c, profileQuery, userID, req.Username, "")
	if err != nil {
		log.Printf("ERROR: Failed to create profile for user %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(userID), req.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for user %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User registered successfully, id=%d, username='%s', email='%s'", userID, req.Username, req.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := database.Pool.QueryRow(c, query, req.Email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		log.Printf("ERROR: Database error during login for email '%s': %v", req.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Printf("WARN: Failed login attempt for email '%s'", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(user.ID), user.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for user %d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User logged in successfully, id=%d, email='%s'", user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
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
		Code     string `json:"code" binding:"required,len=6"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	codeInfo, exists := h.codes[req.Email]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no verification code found"})
		return
	}

	if time.Since(codeInfo.createdAt) > time.Minute*15 {
		delete(h.codes, req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "verification code expired"})
		return
	}

	if codeInfo.code != req.Code {
		log.Printf("WARN: Invalid verification code attempt for email '%s'", req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR: Password hashing failed during email verification for '%s': %v", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	var userID int
	err = database.Pool.QueryRow(c, query, req.Username, req.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		log.Printf("ERROR: Failed to create user during email verification, username='%s': %v", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	profileQuery := `
		INSERT INTO profiles (user_id, name, surname, hackathon_id)
		VALUES ($1, $2, $3, NULL)`

	_, err = database.Pool.Exec(c, profileQuery, userID, req.Username, "")
	if err != nil {
		log.Printf("ERROR: Failed to create profile during email verification for user %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(userID), req.Email)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens during email verification for user %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	delete(h.codes, req.Email)

	log.Printf("INFO: User registered via email verification, id=%d, username='%s', email='%s'", userID, req.Username, req.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// TelegramLogin: возвращает ссылку на авторизацию через Telegram
func (h *AuthHandler) TelegramLogin(c *gin.Context) {
	authURL := h.telegramService.GenerateAuthURL()
	c.JSON(http.StatusOK, gin.H{
		"auth_url": authURL,
	})
}

// VerifyTelegram: проверяет данные, полученные от Telegram Login Widget
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

	accessToken, refreshToken, err := utils.GenerateTokenPair(1, username)
	if err != nil {
		log.Printf("ERROR: Failed to generate tokens for Telegram auth: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	log.Printf("INFO: User logged in via Telegram, telegram_id='%s', username='%s'", telegramID, username)
	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"telegram_id":   telegramID,
		"username":      username,
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Refresh token endpoint",
	})
}

func (h *AuthHandler) CheckToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Check token endpoint",
	})
}

func (h *AuthHandler) DeleteAccount(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete account endpoint",
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout endpoint",
	})
}
