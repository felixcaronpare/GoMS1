package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique;not null;size:255"`
	Password string `gorm:"size:255"`
}