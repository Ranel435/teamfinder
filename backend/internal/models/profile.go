package models

type Profile struct {
	ID             int      `json:"id"`
	UserID         int      `json:"user_id"`
	HackathonID    int      `json:"hackathon_id"`
	Name           string   `json:"name"`
	Surname        string   `json:"surname"`
	AcademicGroup  string   `json:"academic_group"`
	TelegramHandle string   `json:"telegram_handle"`
	DesiredRole    string   `json:"desired_role"`
	Skills         []string `json:"skills"`
	AboutMe        string   `json:"about_me"`
	Achievements   []string `json:"achievements"`
	Status         string   `json:"status"`
}
