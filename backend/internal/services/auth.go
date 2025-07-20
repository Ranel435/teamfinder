package services

import (
	"context"
	"database/sql"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
	"teamfinder/backend/internal/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	var createdAt time.Time
	err = database.Pool.QueryRow(context.Background(), query, username, email, string(hashedPassword)).
		Scan(&user.ID, &createdAt)

	return user, err
}

func (s *AuthService) CreateDefaultProfile(userID int, name string) error {
	query := `
		INSERT INTO profiles (user_id, name, surname, hackathon_id, status)
		VALUES ($1, $2, $3, NULL, 'inactive')`

	_, err := database.Pool.Exec(context.Background(), query, userID, name, "")
	return err
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`

	err := database.Pool.QueryRow(context.Background(), query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	return user, err
}

func (s *AuthService) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash FROM users WHERE id = $1`

	err := database.Pool.QueryRow(context.Background(), query, userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (s *AuthService) ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *AuthService) GenerateTokens(userID int, email string) (string, string, error) {
	return utils.GenerateTokenPair(uint(userID), email)
}

func (s *AuthService) DeleteUser(userID int) error {
	// Сначала удаляем все профили пользователя (CASCADE должен это сделать, но для уверенности)
	_, err := database.Pool.Exec(context.Background(),
		"DELETE FROM profiles WHERE user_id = $1", userID)
	if err != nil {
		return err
	}

	// Удаляем пользователя
	_, err = database.Pool.Exec(context.Background(),
		"DELETE FROM users WHERE id = $1", userID)
	return err
}
