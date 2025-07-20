package services

import (
	"context"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
)

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfilesByHackathonID(hackathonID int) ([]models.Profile, error) {
	query := `
		SELECT id, user_id, hackathon_id, name, surname, academic_group, 
			   telegram_handle, desired_role, skills, about_me, achievements, status 
		FROM profiles
		WHERE hackathon_id = $1 AND status = 'active'
		ORDER BY id DESC`

	rows, err := database.Pool.Query(context.Background(), query, hackathonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var p models.Profile
		err := rows.Scan(&p.ID, &p.UserID, &p.HackathonID, &p.Name, &p.Surname,
			&p.AcademicGroup, &p.TelegramHandle, &p.DesiredRole, &p.Skills,
			&p.AboutMe, &p.Achievements, &p.Status)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func (s *ProfileService) CreateProfile(profile *models.Profile) (int, error) {
	query := `
		INSERT INTO profiles (user_id, hackathon_id, name, surname, academic_group, telegram_handle, 
			desired_role, skills, about_me, achievements, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

	var id int
	err := database.Pool.QueryRow(context.Background(), query,
		profile.UserID, profile.HackathonID, profile.Name, profile.Surname, profile.AcademicGroup,
		profile.TelegramHandle, profile.DesiredRole, profile.Skills,
		profile.AboutMe, profile.Achievements, profile.Status).Scan(&id)

	return id, err
}

func (s *ProfileService) GetProfileByID(hackathonID, profileID int) (*models.Profile, error) {
	query := `
		SELECT id, user_id, hackathon_id, name, surname, academic_group,
			   telegram_handle, desired_role, skills, about_me, achievements, status
		FROM profiles 
		WHERE id = $1 AND hackathon_id = $2`

	var profile models.Profile
	err := database.Pool.QueryRow(context.Background(), query, profileID, hackathonID).Scan(
		&profile.ID, &profile.UserID, &profile.HackathonID, &profile.Name, &profile.Surname,
		&profile.AcademicGroup, &profile.TelegramHandle, &profile.DesiredRole,
		&profile.Skills, &profile.AboutMe, &profile.Achievements, &profile.Status)

	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *ProfileService) UpdateProfile(hackathonID, profileID int, profile *models.Profile) error {
	query := `
		UPDATE profiles SET 
			name = $1, surname = $2, academic_group = $3, telegram_handle = $4,
			desired_role = $5, skills = $6, about_me = $7, achievements = $8,
			status = $9, updated_at = CURRENT_TIMESTAMP
		WHERE id = $10 AND hackathon_id = $11`

	_, err := database.Pool.Exec(context.Background(), query,
		profile.Name, profile.Surname, profile.AcademicGroup, profile.TelegramHandle,
		profile.DesiredRole, profile.Skills, profile.AboutMe, profile.Achievements,
		profile.Status, profileID, hackathonID)

	return err
}

func (s *ProfileService) DeleteProfile(hackathonID, profileID int) error {
	query := `DELETE FROM profiles WHERE id = $1 AND hackathon_id = $2`
	_, err := database.Pool.Exec(context.Background(), query, profileID, hackathonID)
	return err
}
