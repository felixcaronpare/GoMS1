package repo

import (
	"github.com/felixcaronpare/GoMS1/internal/models"
	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.repoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(account models.Account) error {
	err := repo.db.Create(&account).Error

	return err
}
