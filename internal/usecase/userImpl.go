package usecase

import (
	"github.com/adilsonmenechini/goapi/internal/entity"
	"github.com/adilsonmenechini/goapi/internal/presenter"
	"github.com/adilsonmenechini/goapi/internal/repository"
)

type userservice struct {
	repository repository.UserRepository
}

func UserNewService(r repository.UserRepository) UserService {
	return &userservice{repository: r}
}

func (s *userservice) Creater(user *entity.User) (*entity.User, error) {
	user, err := s.repository.Insert(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userservice) Update(user *entity.User) (*entity.User, error) {
	user, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userservice) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *userservice) Fetch(id string) (entity.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userservice) Fetchs() ([]presenter.UserResponse, error) {
	users, err := s.repository.FindByAll()
	if err != nil {
		return nil, err
	}
	response := []presenter.UserResponse{}
	for _, usr := range users {
		responses := presenter.UserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
			Token: usr.Token,
		}
		response = append(response, responses)
	}
	return response, nil
}
