package repo

import (
	"github.com/felixcaronpare/GoMS1/internal/models"
	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

//=========== CRUD ===========

func (repo *Repo) Create(account models.Account) (models.Account, error) {
	err := repo.db.Create(&account).Error

	return account, err
}

//=========== UTILS ===========

func (repo *Repo) GetByEmail(email string) (models.Account, error){
	var account models.Account
	err := repo.db.Where("email = ?", email).First(&account).Error
  
	return account, err
}