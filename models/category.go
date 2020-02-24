package models

// Category is the model for categories
type Category struct {
	Id   uint   `gorm:"primary_key"json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);not null;unique"json:"name,omitempty"`
}
