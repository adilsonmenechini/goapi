package repository

import (
	"github.com/adilsonmenechini/goapi/internal/entity"
	"gorm.io/gorm"
)

type Reader interface {
	FindByAll() ([]entity.User, error)
	FindEmail(email string) (entity.User, error)
	FindByID(id string) (entity.User, error)
}

type Writer interface {
	Insert(user *entity.User) (*entity.User, error)
	Delete(id string) error
	Save(user *entity.User) (*entity.User, error)
}

type UserRepository interface {
	Reader
	Writer
}

type userrepository struct {
	DB *gorm.DB
}
