package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"sync"
	"time"
	// "github.com/joho/godotenv"
)

type EmailService struct {
	from     string
	password string
	host     string
	port     string
	lastSent map[string]time.Time
	mutex    sync.RWMutex
}

func NewEmailService() *EmailService {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	// }

	return &EmailService{
		from:     os.Getenv("EMAIL_LOGIN"),
		password: os.Getenv("EMAIL_PASSWORD"),
		host:     "smtp.mail.ru",
		port:     "465",
		lastSent: make(map[string]time.Time),
	}
}

func (s *EmailService) GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (s *EmailService) SendVerificationCode(to, code string) error {
	log.Printf("Starting to send verification code. SMTP settings - Host: %s, Port: %s, From: %s", s.host, s.port, s.from)

	// Проверяем, прошло ли 5 минут с последней отправки
	s.mutex.RLock()
	lastTime, exists := s.lastSent[to]
	s.mutex.RUnlock()

	if exists {
		timeSinceLastEmail := time.Since(lastTime)
		if timeSinceLastEmail < 5*time.Minute {
			remainingTime := 5*time.Minute - timeSinceLastEmail
			return fmt.Errorf("please wait %v before requesting a new code", remainingTime.Round(time.Second))
		}
	}

	// Настройка TLS конфигурации
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Добавлено для отладки
		ServerName:         s.host,
	}

	// Подключение к серверу с TLS
	log.Printf("Attempting to connect to SMTP server...")
	conn, err := tls.Dial("tcp", s.host+":"+s.port, tlsConfig)
	if err != nil {
		log.Fatalf("failed to create TLS connection: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.host)
	if err != nil {
		log.Fatalf("failed to create SMTP client: %v", err)
	}
	defer client.Close()

	// Аутентификация
	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	if err := client.Auth(auth); err != nil {
		log.Fatalf("failed to authenticate: %v", err)
	}

	// Отправка письма
	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Verification Code\r\n"+
		"\r\n"+
		"Your verification code is: %s\r\n", s.from, to, code))

	if err := client.Mail(s.from); err != nil {
		return err
	}
	if err := client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	// После успешной отправки обновляем время
	s.mutex.Lock()
	s.lastSent[to] = time.Now()
	s.mutex.Unlock()

	return client.Quit()
}
