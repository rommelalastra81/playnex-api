package models

type Users struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
