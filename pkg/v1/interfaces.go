package v1

import "github.com/felixcaronpare/GoMS1/internal/models"

type RepoInterface interface {
	Create(models.Account) error
}

type UsecaseInterface interface {
	Create(models.Account) error
}
