package hackathon

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

func (r *Repository) CreateHackathon(hackathon *models.Hackathon) error {
	query := `
		INSERT INTO hackathons (name, description, start_date, end_date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`

	err := database.Pool.QueryRow(
		context.Background(),
		query,
		hackathon.Name,
		hackathon.Description,
		hackathon.StartDate,
		hackathon.EndDate,
	).Scan(&hackathon.ID)

	return err
}

func (r *Repository) GetHackathonByID(id int) (*models.Hackathon, error) {
	hackathon := &models.Hackathon{}
	query := `
		SELECT id, name, description, start_date, end_date 
		FROM hackathons 
		WHERE id = $1
	`

	err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&hackathon.ID,
		&hackathon.Name,
		&hackathon.Description,
		&hackathon.StartDate,
		&hackathon.EndDate,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return hackathon, err
}

func (r *Repository) GetAllHackathons() ([]*models.Hackathon, error) {
	query := `
		SELECT id, name, description, start_date, end_date 
		FROM hackathons 
		ORDER BY start_date DESC
	`

	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hackathons []*models.Hackathon
	for rows.Next() {
		hackathon := &models.Hackathon{}
		err := rows.Scan(
			&hackathon.ID,
			&hackathon.Name,
			&hackathon.Description,
			&hackathon.StartDate,
			&hackathon.EndDate,
		)
		if err != nil {
			return nil, err
		}
		hackathons = append(hackathons, hackathon)
	}

	return hackathons, nil
}

func (r *Repository) UpdateHackathon(hackathon *models.Hackathon) error {
	query := `
		UPDATE hackathons SET 
			name = $1, description = $2, start_date = $3, end_date = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
	`

	_, err := database.Pool.Exec(
		context.Background(),
		query,
		hackathon.Name,
		hackathon.Description,
		hackathon.StartDate,
		hackathon.EndDate,
		hackathon.ID,
	)

	return err
}

func (r *Repository) DeleteHackathon(id int) error {
	query := `DELETE FROM hackathons WHERE id = $1`
	_, err := database.Pool.Exec(context.Background(), query, id)
	return err
}
