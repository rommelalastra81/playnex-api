package models

import "time"

type Clubs struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	LogoUrl       string    `json:"logo_url"`
	CoverImageUrl string    `json:"cover_image_url"`
	CreatorId     uint      `json:"creator_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
