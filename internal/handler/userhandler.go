package handler

import (
	"net/http"

	"github.com/adilsonmenechini/goapi/internal/entity"
	"github.com/adilsonmenechini/goapi/internal/presenter"
	"github.com/adilsonmenechini/goapi/internal/usecase"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func UserNewHandler(s usecase.UserService) UserHandler {
	return &userhandler{service: s}
}

// @Summary Creater user
// @Description get string creater user
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user body presenter.UserRequest true "Creater user"
// @Router /api/user/ [post]
func (h userhandler) AddUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user entity.User
		err := c.BodyParser(&user)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		result, err := h.service.Creater(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// @Summary Update user
// @Description get string Update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user body presenter.UserRequest true "Update user"
// @Param id path string true "Update user"
// @Router /api/user/{id} [put]
func (h userhandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userreq presenter.UserRequest
		err := c.BodyParser(&userreq)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		user, err := h.service.Fetch(id)
		if err != nil {
			return err
		}
		user.Name = userreq.Name
		user.Email = userreq.Email
		user.Password = userreq.Password

		result, err := h.service.Update(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}

}

// @Summary Delete user
// @Description get string Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Delete user"
// @Router /api/user/{id} [delete]
func (h userhandler) RemoveUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		err := h.service.Delete(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "delete successfully",
			"err":    nil,
		})
	}

}

// @Summary Get user
// @Description get string Get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Get user"
// @Router /api/user/{id} [get]
func (h userhandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		user, err := h.service.Fetch(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(&user))
	}
}

// @Summary Alls user
// @Description get string Alls user
// @Tags user
// @Accept  json
// @Produce  json
// @Router /api/user/ [get]
func (h userhandler) GetUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := h.service.Fetchs()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersFindResponse(users))
	}
}
