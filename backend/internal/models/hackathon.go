package models

type Hackathon struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"AI Innovation Hackathon 2024"`
	Description string `json:"description" example:"48-hour hackathon focused on artificial intelligence and machine learning solutions"`
	StartDate   string `json:"start_date" example:"2024-03-15T09:00:00Z"`
	EndDate     string `json:"end_date" example:"2024-03-17T18:00:00Z"`
}
