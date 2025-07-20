package auth

import (
	"context"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`

	err := database.Pool.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID)

	return err
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash 
		FROM users 
		WHERE email = $1
	`

	err := database.Pool.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash 
		FROM users 
		WHERE username = $1
	`

	err := database.Pool.QueryRow(context.Background(), query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash 
		FROM users 
		WHERE id = $1
	`

	err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return user, err
}
