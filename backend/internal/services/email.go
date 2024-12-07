package services

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	// "github.com/joho/godotenv"
)

type EmailService struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailService() *EmailService {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	// }

	return &EmailService{
		from:     os.Getenv("EMAIL_LOGIN"),
		password: os.Getenv("EMAIL_PASSWORD"),
		host:     "smtp.gmail.com",
		port:     "587",
	}
}

func (s *EmailService) GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (s *EmailService) SendVerificationCode(to, code string) error {
	auth := smtp.PlainAuth("", s.from, s.password, s.host)

	msg := []byte(fmt.Sprintf("Subject: Verification Code\r\n\r\nYour verification code is: %s", code))

	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	err := smtp.SendMail(addr, auth, s.from, []string{to}, msg)
	if err != nil {
		log.Printf("Ошибка отправки email: %v", err)
	}
	return err
}
