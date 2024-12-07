package models

type Profile struct {
	ID          int      `json:"id"`
	HackathonID int      `json:"hackathon_id"`
	UserID      int      `json:"user_id"`
	Name        string   `json:"name"`
	Skills      []string `json:"skills"`
	Description string   `json:"description"`
	// Add other fields as needed
}
