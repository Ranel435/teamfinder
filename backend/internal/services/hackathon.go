package services

import (
	"context"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
)

type HackathonService struct {
	// Add dependencies like database connection here
}

func NewHackathonService() *HackathonService {
	return &HackathonService{}
}

func (s *HackathonService) GetAllHackathons() ([]models.Hackathon, error) {
	query := `SELECT id, name, description, start_date, end_date FROM hackathons`
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hackathons []models.Hackathon
	for rows.Next() {
		var h models.Hackathon
		err := rows.Scan(&h.ID, &h.Name, &h.Description, &h.StartDate, &h.EndDate)
		if err != nil {
			return nil, err
		}
		hackathons = append(hackathons, h)
	}
	return hackathons, nil
}

func (s *HackathonService) GetHackathonByID(id int) (*models.Hackathon, error) {
	return nil, nil // Temporary nil return
}

func (s *HackathonService) CreateHackathon(hackathon *models.Hackathon) (int, error) {
	return 0, nil // Temporary zero return
}

func (s *HackathonService) UpdateHackathon(id int, hackathon *models.Hackathon) error {
	return nil // Temporary nil return
}

func (s *HackathonService) DeleteHackathon(id int) error {
	return nil // Temporary nil return
}
