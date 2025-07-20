package profile

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

func (r *Repository) CreateProfile(profile *models.Profile) error {
	query := `
		INSERT INTO profiles (
			user_id, hackathon_id, name, surname, academic_group, 
			telegram_handle, desired_role, skills, about_me, achievements, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING id
	`

	err := database.Pool.QueryRow(
		context.Background(),
		query,
		profile.UserID,
		profile.HackathonID,
		profile.Name,
		profile.Surname,
		profile.AcademicGroup,
		profile.TelegramHandle,
		profile.DesiredRole,
		profile.Skills,
		profile.AboutMe,
		profile.Achievements,
		profile.Status,
	).Scan(&profile.ID)

	return err
}

func (r *Repository) GetProfileByUserAndHackathon(userID, hackathonID int) (*models.Profile, error) {
	profile := &models.Profile{}
	query := `
		SELECT id, user_id, hackathon_id, name, surname, academic_group,
			   telegram_handle, desired_role, skills, about_me, achievements, status
		FROM profiles 
		WHERE user_id = $1 AND hackathon_id = $2
	`

	err := database.Pool.QueryRow(context.Background(), query, userID, hackathonID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.HackathonID,
		&profile.Name,
		&profile.Surname,
		&profile.AcademicGroup,
		&profile.TelegramHandle,
		&profile.DesiredRole,
		&profile.Skills,
		&profile.AboutMe,
		&profile.Achievements,
		&profile.Status,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return profile, err
}

func (r *Repository) GetProfilesByHackathon(hackathonID int) ([]*models.Profile, error) {
	query := `
		SELECT id, user_id, hackathon_id, name, surname, academic_group,
			   telegram_handle, desired_role, skills, about_me, achievements, status
		FROM profiles 
		WHERE hackathon_id = $1 AND status = 'active'
		ORDER BY created_at DESC
	`

	rows, err := database.Pool.Query(context.Background(), query, hackathonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		profile := &models.Profile{}
		err := rows.Scan(
			&profile.ID,
			&profile.UserID,
			&profile.HackathonID,
			&profile.Name,
			&profile.Surname,
			&profile.AcademicGroup,
			&profile.TelegramHandle,
			&profile.DesiredRole,
			&profile.Skills,
			&profile.AboutMe,
			&profile.Achievements,
			&profile.Status,
		)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (r *Repository) UpdateProfile(profile *models.Profile) error {
	query := `
		UPDATE profiles SET 
			name = $1, surname = $2, academic_group = $3, telegram_handle = $4,
			desired_role = $5, skills = $6, about_me = $7, achievements = $8,
			status = $9, updated_at = CURRENT_TIMESTAMP
		WHERE id = $10
	`

	_, err := database.Pool.Exec(
		context.Background(),
		query,
		profile.Name,
		profile.Surname,
		profile.AcademicGroup,
		profile.TelegramHandle,
		profile.DesiredRole,
		profile.Skills,
		profile.AboutMe,
		profile.Achievements,
		profile.Status,
		profile.ID,
	)

	return err
}
