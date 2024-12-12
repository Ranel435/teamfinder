package services

import (
	"context"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
)

type ProfileService struct {
	// Add dependencies like database connection here
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfilesByHackathonID(hackathonID int) ([]models.Profile, error) {
	query := `
		SELECT p.id, p.user_id, p.name, p.academic_group, p.telegram_handle, 
			   p.desired_role, p.skills, p.about_me, p.achievements, p.status 
		FROM profiles p
		JOIN hackathon_profiles hp ON p.id = hp.profile_id
		WHERE hp.hackathon_id = $1`

	rows, err := database.Pool.Query(context.Background(), query, hackathonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var p models.Profile
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.AcademicGroup,
			&p.TelegramHandle, &p.DesiredRole, &p.Skills,
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
		INSERT INTO profiles (user_id, name, surname, academic_group, telegram_handle, 
			desired_role, skills, about_me, achievements, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`

	var id int
	err := database.Pool.QueryRow(context.Background(), query,
		profile.UserID, profile.Name, profile.Surname, profile.AcademicGroup,
		profile.TelegramHandle, profile.DesiredRole, profile.Skills,
		profile.AboutMe, profile.Achievements, profile.Status).Scan(&id)

	return id, err
}

func (s *ProfileService) GetProfileByID(hackathonID, profileID int) (*models.Profile, error) {
	return nil, nil
}

func (s *ProfileService) UpdateProfile(hackathonID, profileID int, profile *models.Profile) error {
	return nil
}

func (s *ProfileService) DeleteProfile(hackathonID, profileID int) error {
	return nil
}
