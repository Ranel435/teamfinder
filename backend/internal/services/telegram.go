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
	"strings"
	// "github.com/joho/godotenv"
)

type TelegramService struct {
	botToken string
}

func NewTelegramService() *TelegramService {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	// }

	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		log.Fatalf("Токен не найден в .env файле")
	}

	return &TelegramService{
		botToken: token,
	}
}

func (s *TelegramService) GenerateAuthURL() string {
	baseURL := "https://telegram.org/oauth/authorize"
	params := url.Values{}
	params.Add("bot_id", os.Getenv("TG_BOT_ID"))
	//THERE WILL BE OUR LINK TO DOMEN
	params.Add("origin", "https://your-website.com")
	params.Add("return_to", "https://your-website.com/auth/callback")

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func (s *TelegramService) ValidateAuthData(data map[string]string) bool {
	// Check if hash is present
	hash, ok := data["hash"]
	if !ok {
		return false
	}
	delete(data, "hash")

	// Sort all key-value pairs
	var dataCheckString []string
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		dataCheckString = append(dataCheckString, fmt.Sprintf("%s=%s", k, data[k]))
	}

	// Create data check string
	dataStr := strings.Join(dataCheckString, "\n")

	// Calculate secret key
	secretKey := sha256.Sum256([]byte(s.botToken))

	// Calculate hash
	mac := hmac.New(sha256.New, secretKey[:])
	mac.Write([]byte(dataStr))
	expectedHash := hex.EncodeToString(mac.Sum(nil))

	return hash == expectedHash
}
