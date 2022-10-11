package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/adilsonmenechini/goapi/internal/entity"
	"github.com/adilsonmenechini/goapi/internal/pkg/db"
	"github.com/adilsonmenechini/goapi/internal/repository"
	"github.com/adilsonmenechini/goapi/internal/usecase"
	"github.com/bxcodec/faker/v3"
)

func TestUserUseCase_Create(t *testing.T) {
	var usr entity.User
	usr.Name = faker.Name()
	usr.Email = faker.Email()
	usr.Password = "123456asddfg"

	db := db.ConnectMySQL()
	repo := repository.UserNewRepo(db)
	serv := usecase.UserNewService(repo)
	user, err := serv.Creater(&usr)
	if err != nil {
		log.Fatalf("Errro ## %v", err)
	}
	fmt.Println(user)
}

func TestUserUseCase_FindID(t *testing.T) {
	db := db.ConnectMySQL()
	repo := repository.UserNewRepo(db)
	serv := usecase.UserNewService(repo)
	result, err := serv.Fetch("f18e5809-6ad8-4575-bf76-83810a037fdc")
	if err != nil {
		log.Println("Error")
		return
	}
	fmt.Println(result)
}

func TestUserUseCase_FindEmail(t *testing.T) {
	db := db.ConnectMySQL()
	repo := repository.UserNewRepo(db)
	result, err := repo.FindEmail("EvhgYPB@aASeEVl.ru")
	if err != nil {
		log.Println("Error")
		return
	}
	fmt.Println(result)
}

func TestUserUseCase_CreateP(t *testing.T) {
	var user entity.User
	user.Name = faker.Name()
	user.Email = faker.Email()
	user.Password = "123456asddfg"

	db := db.ConnectPSQL()
	repo := repository.UserNewRepo(db)
	serv := usecase.UserNewService(repo)
	result, err := serv.Creater(&user)
	if err != nil {
		log.Fatalf("Errro ## %v", err)
	}
	fmt.Println(result)
}

func TestUserUseCase_FindEmailP(t *testing.T) {
	db := db.ConnectPSQL()
	repo := repository.UserNewRepo(db)
	result, err := repo.FindEmail("adilsonpostman05@gmail.com")
	if err != nil {
		log.Println("Error")
		return
	}
	fmt.Println(result)
}
