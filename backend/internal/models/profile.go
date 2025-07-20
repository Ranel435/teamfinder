package models

type Profile struct {
	ID             int      `json:"id" example:"1"`
	UserID         int      `json:"user_id" example:"5"`
	HackathonID    int      `json:"hackathon_id" example:"1"`
	Name           string   `json:"name" example:"John"`
	Surname        string   `json:"surname" example:"Doe"`
	AcademicGroup  string   `json:"academic_group" example:"CS-401"`
	TelegramHandle string   `json:"telegram_handle" example:"@johndoe"`
	DesiredRole    string   `json:"desired_role" example:"Backend Developer"`
	Skills         []string `json:"skills" example:"Go,PostgreSQL,Docker,REST API"`
	AboutMe        string   `json:"about_me" example:"Passionate backend developer with 2 years of experience in Go and microservices"`
	Achievements   []string `json:"achievements" example:"1st place in AI Hackathon 2023,Completed 10+ projects"`
	Status         string   `json:"status" example:"active"`
}
