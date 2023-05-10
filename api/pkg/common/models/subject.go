package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  int    `json:"catId"`
}

