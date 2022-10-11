package routes

import (
	"github.com/adilsonmenechini/goapi/api/handler"
	"github.com/adilsonmenechini/goapi/internal/pkg/db"
	"github.com/adilsonmenechini/goapi/internal/repository"
	"github.com/adilsonmenechini/goapi/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(internal fiber.Router) {
	db := db.ConnectPSQL()
	r := repository.UserNewRepo(db)
	serv := usecase.UserNewService(r)
	hdr := handler.UserNewHandler(serv)
	user := internal.Group("/user")
	user.Get("/", hdr.GetUsers())
	user.Post("/", hdr.AddUser())
	user.Put("/:id", hdr.UpdateUser())
	user.Get("/:id", hdr.GetUser())
	user.Delete("/:id", hdr.RemoveUser())

}
