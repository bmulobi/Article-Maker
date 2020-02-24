// Package models contains the definitions for the database schema
package models

import "time"

// Article is the model for articles
type Article struct {
	Id uint `gorm:"primary_key"json:"id,omitempty"`
	Title string `gorm:"type:varchar(100);not null;unique"json:"title,omitempty"`
	Body string `gorm:"type:text;not null"json:"body,omitempty"`
	PublisherID uint `gorm:"not null"json:"publisher_id,omitempty"`
	Publisher Publisher `json:"publisher,omitempty"`
	CategoryID uint `gorm:"not null"json:"category_id,omitempty"`
	Category Category `json:"category,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	PublishedAt time.Time `gorm:"default:null"json:"published_at,omitempty"`
}
