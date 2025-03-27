package usecase

import (
	"errors"

	"github.com/felixcaronpare/GoMS1/internal/models"
	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	// accountRepo "github.com/felixcaronpare/GoMS1/pkg/v1/repository"
	"gorm.io/gorm"
)

type Usecase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UsecaseInterface {
	return &Usecase{repo}
}

//=========== CRUD ===========
//Create
func (uc *Usecase) Create(account models.Account) (models.Account, error) {
	if _, err := uc.repo.GetByEmail(account.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Account{}, errors.New("The provided email is already associated with another account.")
	}

	return uc.repo.Create(account)
}

//Read
func (uc *UseCase) Get(id string) (models.User, error){
	var user models.User
	var err error
  
	if user, err = uc.repo.Get(id); err != nil{
	  if errors.Is(err, gorm.ErrRecordNotFound){
		return models.User{}, errors.New("no such user with the id supplied")
	  }
	  // handle the error properly as the error was something else then record not found,
	  // and return to the caller of this function
	  return models.User{}, err
	}
  
	return user, nil
  }
  
  // Update
  func (uc *UseCase) Update(updateUser models.User) error{
	var user models.User
	var err error
	// check if user exists
	if user, err = uc.Get(string(updateUser.ID)); err != nil {
	  return err
	}
  
	// check if only name is going to change,
	// as the email cannot be changed
	if user.Email != updateUser.Email {
	  return errors.New("email cannot be changed")
	}
	
	err = uc.repo.Update(updateUser)
	if err != nil {
	  // handle the error properly as the error might be something worth to debug
	}
  
	return nil
  }
  
  // Delete
  func (uc *UseCase) Delete(id string) error{
	var err error
	// check if user exists
	if _, err = uc.Get(id); err != nil {
	  return err
	}
  
	err = uc.repo.Delete(id)
	if err != nil {
	  // handle the error as it might be something worth to debug
	  return err
	}
  
	return nil
  }