package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	email    string `gorm:"unique;not null"`
	password string
}
