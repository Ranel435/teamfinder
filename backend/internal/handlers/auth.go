package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"
	"teamfinder/backend/internal/db"
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

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Create user in database
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	var userID int
	err = db.Pool.QueryRow(c, query, req.Username, req.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Create empty profile for user
	profileQuery := `
		INSERT INTO profiles (user_id, name)
		VALUES ($1, $2)`

	_, err = db.Pool.Exec(c, profileQuery, userID, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(userID), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

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

	// Get user from database
	var user models.User
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := db.Pool.QueryRow(c, query, req.Email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send verification code",
		})
		return
	}

	h.codes[req.Email] = codeInfo{
		code:      code,
		createdAt: time.Now(),
	}

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

	// Добавляем логирование входящих данных
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Printf("Received request body: %s\n", string(body))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fmt.Printf("Received data: email=%s, code=%s, username=%s, password=%s\n",
			req.Email, req.Code, req.Username, req.Password)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Create user in database
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	var userID int
	err = db.Pool.QueryRow(c, query, req.Username, req.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Create empty profile for user
	profileQuery := `
		INSERT INTO profiles (user_id, name)
		VALUES ($1, $2)`

	_, err = db.Pool.Exec(c, profileQuery, userID, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(userID), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	delete(h.codes, req.Email)

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

	// Извлекаем параметры из тела запроса
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Проверяем подлинность данных с помощью TelegramService
	if !h.telegramService.ValidateAuthData(req) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Telegram data"})
		return
	}

	// Извлекаем данные пользователя
	telegramID := req["id"]
	username := req["username"]

	// Здесь надо будет добавить логику проверки или регистрации пользователя в БД
	// Например: findOrCreateUser(telegramID, username)

	// Создаем токены
	accessToken, refreshToken, err := utils.GenerateTokenPair(1, username) // Реальный ID пользователя заменить на 1
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Возвращаем токены клиенту
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
