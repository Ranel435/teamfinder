package services

import (
	"context"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/models"
)

type HackathonService struct{}

func NewHackathonService() *HackathonService {
	return &HackathonService{}
}

func (s *HackathonService) GetAllHackathons() ([]models.Hackathon, error) {
	query := `SELECT id, name, description, 
		start_date::text as start_date, 
		end_date::text as end_date 
		FROM hackathons ORDER BY id DESC`
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
	var hackathon models.Hackathon
	query := `SELECT id, name, description, 
		start_date::text as start_date, 
		end_date::text as end_date 
		FROM hackathons WHERE id = $1`
	err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&hackathon.ID, &hackathon.Name, &hackathon.Description,
		&hackathon.StartDate, &hackathon.EndDate)
	if err != nil {
		return nil, err
	}
	return &hackathon, nil
}

func (s *HackathonService) CreateHackathon(hackathon *models.Hackathon) (int, error) {
	query := `
		INSERT INTO hackathons (name, description, start_date, end_date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`

	var id int
	err := database.Pool.QueryRow(context.Background(), query,
		hackathon.Name, hackathon.Description, hackathon.StartDate, hackathon.EndDate).Scan(&id)
	return id, err
}

func (s *HackathonService) UpdateHackathon(id int, hackathon *models.Hackathon) error {
	query := `
		UPDATE hackathons SET 
			name = $1, description = $2, start_date = $3, end_date = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $5`

	_, err := database.Pool.Exec(context.Background(), query,
		hackathon.Name, hackathon.Description, hackathon.StartDate, hackathon.EndDate, id)
	return err
}

func (s *HackathonService) DeleteHackathon(id int) error {
	query := `DELETE FROM hackathons WHERE id = $1`
	_, err := database.Pool.Exec(context.Background(), query, id)
	return err
}
