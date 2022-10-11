package usecase

import (
	"github.com/adilsonmenechini/goapi/internal/entity"
	"github.com/adilsonmenechini/goapi/internal/presenter"
)

type Writer interface {
	Creater(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
}

type Reader interface {
	Fetch(id string) (entity.User, error)
	Fetchs() ([]presenter.UserResponse, error)
}

type UserService interface {
	Writer
	Reader
}
