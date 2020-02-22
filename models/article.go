package models

import "time"

// Article is the model for articles
type Article struct {
	Id uint `gorm:"primary_key";json:"id,omitempty"`
	Title string `gorm:"type:varchar(100);not null"`
	Body string `gorm:"type:text;not null"`
	PublisherID uint `gorm:"default:NULL";json:"publisher_id,omitempty"`
	Publisher Publisher `json:"publisher,omitempty"`
	CategoryID uint `gorm:"default:NULL";json:"category_id,omitempty"`
	Category Category `json:"category,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	PublishedAt time.Time `gorm:"default:NULL";json:"published_at,omitempty"`
}
