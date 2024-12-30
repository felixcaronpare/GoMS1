package usecase

import (
	"errors"

	"github.com/felixcaronpare/GoMS1/internals/models"
	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	accountRepo "github.com/felixcaronpare/GomS1/pkg/v1/repository"
	"gorm.io/gorm"
)

type Usecase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UsecaseInterface {
	return &Usecase{repo}
}

