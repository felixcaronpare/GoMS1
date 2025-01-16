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
// Create
func (repo *Repo) Create(account models.Account) (models.Account, error) {
	err := repo.db.Create(&account).Error

	return account, err
}

// Get
func (repo *Repo) Get(id string) (models.Account, error){
	var account models.Account
	err := repo.db.Where("id = ?", id).First(&account).Error
  
	return account, err
  }
  
// Update
func (repo *Repo) Update(account models.Account) error{
var dbAccount models.Account
if err := repo.db.Where("id = ?", account.ID).First(&dbAccount).Error; err != nil {
	return err
}
dbAccount.Name = account.Name
err := repo.db.Save(dbAccount).Error

return err
}

// Delete
func (repo *Repo) Delete(id string) error{
err := repo.db.Where("id = ?", id).Delete(&models.Account{}).Error

return err
}

//=========== UTILS ===========

func (repo *Repo) GetByEmail(email string) (models.Account, error){
	var account models.Account
	err := repo.db.Where("email = ?", email).First(&account).Error
  
	return account, err
}