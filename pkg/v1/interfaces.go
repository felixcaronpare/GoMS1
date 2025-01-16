package v1

import "github.com/felixcaronpare/GoMS1/internal/models"

type RepoInterface interface {

	Create(models.Account) (models.Account, error)
	Get(id string) (models.Account, error)
	Update(models.Account) (error)
	Delete(id string) (error)

	GetByEmail(email string) (models.Account, error)
}

type UsecaseInterface interface {
	Create(models.Account) (models.Account, error)
    Get(id string) (models.Account, error)
    Update(models.Account) (error)
    Delete(id string) (error)
}
