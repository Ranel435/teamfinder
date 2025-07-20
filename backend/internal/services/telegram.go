package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type TelegramService struct {
	botToken    string
	botID       string
	baseURL     string
	callbackURL string
}

func NewTelegramService() *TelegramService {
	return &TelegramService{
		botToken:    os.Getenv("TG_BOT_TOKEN"),
		botID:       os.Getenv("TG_BOT_ID"),
		baseURL:     "https://teamfinder-hack.ru",
		callbackURL: "https://teamfinder-hack.ru/auth/telegram/callback",
	}
}

func (s *TelegramService) GenerateAuthURL() string {
	if s.botID == "" {
		log.Printf("WARN: TG_BOT_ID not configured")
		return "https://telegram.org/oauth/authorize?bot_id=YOUR_BOT_ID"
	}

	baseURL := "https://telegram.org/oauth/authorize"
	params := url.Values{}
	params.Add("bot_id", s.botID)
	params.Add("origin", s.baseURL)
	params.Add("return_to", s.callbackURL)

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func (s *TelegramService) ValidateAuthData(data map[string]string) bool {
	if s.botToken == "" {
		log.Printf("WARN: TG_BOT_TOKEN not configured, cannot validate Telegram auth")
		return false
	}

	hash, ok := data["hash"]
	if !ok {
		log.Printf("WARN: No hash in Telegram auth data")
		return false
	}

	// Проверка auth_date (не старше 24 часов)
	authDateStr, ok := data["auth_date"]
	if ok {
		authDate, err := strconv.ParseInt(authDateStr, 10, 64)
		if err == nil {
			authTime := time.Unix(authDate, 0)
			if time.Since(authTime) > 24*time.Hour {
				log.Printf("WARN: Telegram auth data too old")
				return false
			}
		}
	}

	// Создаем копию data без hash для проверки
	dataCheck := make(map[string]string)
	for k, v := range data {
		if k != "hash" {
			dataCheck[k] = v
		}
	}

	// Сортируем ключи
	var dataCheckString []string
	keys := make([]string, 0, len(dataCheck))
	for k := range dataCheck {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		dataCheckString = append(dataCheckString, fmt.Sprintf("%s=%s", k, dataCheck[k]))
	}

	dataStr := strings.Join(dataCheckString, "\n")
	secretKey := sha256.Sum256([]byte(s.botToken))
	mac := hmac.New(sha256.New, secretKey[:])
	mac.Write([]byte(dataStr))
	expectedHash := hex.EncodeToString(mac.Sum(nil))

	isValid := hash == expectedHash
	if !isValid {
		log.Printf("WARN: Invalid Telegram hash. Expected: %s, Got: %s", expectedHash, hash)
		log.Printf("DEBUG: Data string: %s", dataStr)
	}

	return isValid
}
