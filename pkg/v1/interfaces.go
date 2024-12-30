package v1

import "github.com/felixcaronpare/GoMS1/internal/models"

type repoInterface interface {
	Create(models.Account) error
}

type usecaseInterface interface {
	Create(models.Account) error
}
