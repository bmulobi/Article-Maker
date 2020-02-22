package models

type Publisher struct {
	Id uint `gorm:"primary_key";json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);not null";json:"name"`
}
