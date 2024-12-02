package handlers

import (
	"net/http"
	"teamfinder/backend/internal/services"
	"teamfinder/backend/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
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

func NewAuthHandler(emailService *services.EmailService, telegramService *services.TelegramService) *AuthHandler {
	return &AuthHandler{
		emailService:    emailService,
		telegramService: telegramService,
		codes:           make(map[string]codeInfo),
	}
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
	h.codes[req.Email] = codeInfo{
		code:      code,
		createdAt: time.Now(),
	}

	if err := h.emailService.SendVerificationCode(req.Email, code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Verification code sent to email.",
	})
}

func (h *AuthHandler) VerifyEmailCode(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required,len=6"`
	}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code"})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := utils.GenerateTokenPair(1, req.Email) // Replace 1 with actual user ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	delete(h.codes, req.Email)

	c.JSON(http.StatusOK, gin.H{
		"token":         accessToken,
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
