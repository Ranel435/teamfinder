package services

import (
	"fmt"
	"math/rand"
	"net/smtp"
)

type EmailService struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailService() *EmailService {
	return &EmailService{
		from:     "your-email@example.com", // Move to env
		password: "your-password",          // Move to env
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
	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}
