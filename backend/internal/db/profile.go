package db

import (
	"context"
	"time"
)

type Profile struct {
	UserID         int      `json:"user_id"`
	Name           string   `json:"name"`
	AcademicGroup  string   `json:"academic_group"`
	TelegramHandle string   `json:"telegram_handle"`
	DesiredRole    string   `json:"desired_role"`
	Skills         []string `json:"skills"`
	AboutMe        string   `json:"about_me"`
	Achievements   []string `json:"achievements"`
	Status         string   `json:"status"`
	Wishlist       []string `json:"wishlist"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// CreateProfile создает профиль
func CreateProfile(ctx context.Context, profile *Profile) error {
	query := `
    INSERT INTO profiles (
        user_id, name, academic_group, telegram_handle, desired_role, 
        skills, about_me, achievements, status, wishlist, created_at, updated_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
    `
	_, err := Pool.Exec(ctx, query,
		profile.UserID, profile.Name, profile.AcademicGroup, profile.TelegramHandle,
		profile.DesiredRole, profile.Skills, profile.AboutMe, profile.Achievements,
		profile.Status, profile.Wishlist)
	return err
}

// GetProfile возвращает профиль по user_id
func GetProfile(ctx context.Context, userID int) (*Profile, error) {
	query := `SELECT user_id, name, academic_group, telegram_handle, desired_role, 
              skills, about_me, achievements, status, wishlist, created_at, updated_at 
              FROM profiles WHERE user_id = $1`

	var profile Profile
	err := Pool.QueryRow(ctx, query, userID).Scan(
		&profile.UserID, &profile.Name, &profile.AcademicGroup, &profile.TelegramHandle,
		&profile.DesiredRole, &profile.Skills, &profile.AboutMe, &profile.Achievements,
		&profile.Status, &profile.Wishlist, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateProfile обновляет профиль
func UpdateProfile(ctx context.Context, profile *Profile) error {
	query := `
    UPDATE profiles SET
        name = $2, academic_group = $3, telegram_handle = $4, desired_role = $5,
        skills = $6, about_me = $7, achievements = $8, status = $9, wishlist = $10,
        updated_at = NOW()
    WHERE user_id = $1
    `
	_, err := Pool.Exec(ctx, query,
		profile.UserID, profile.Name, profile.AcademicGroup, profile.TelegramHandle,
		profile.DesiredRole, profile.Skills, profile.AboutMe, profile.Achievements,
		profile.Status, profile.Wishlist)
	return err
}
