package models

import "gorm.io/gorm"

type Employee struct {
	model gorm.Model
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Email string `json:"email" gorm:"type:varchar(255);not null"`
	Pass  string `json:"password" gorm:"type:varchar(255);not null"`
}
