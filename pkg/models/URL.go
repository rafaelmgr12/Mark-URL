package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	URL      string `json:"url" gorm:"not null" validate:"required,url"`
	ShortURL string `json:"short_url" validate:"required,url"`
	Hits     int    `json:"hits" gorm:"default:0"`
	UserID   string `json:"user_id" gorm:"not null"`
	User     User
}

func UrlValidated(url *URL) error {
	if err := validator.Validate(url); err != nil {
		return err
	}
	return nil
}