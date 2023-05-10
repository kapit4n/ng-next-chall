package models

import "gorm.io/gorm"

type SubjectEvent struct {
	gorm.Model
	Description string `json:"description"`
	SubjectId   int    `json:"subjectId"`
}
