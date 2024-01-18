package handlers

import (
	"EchoApp/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{
		service: s,
	}
}

func (uc *UserController) Hello(c echo.Context) error {
	return c.String(http.StatusOK, uc.service.Hello())
	//uc.service.Hello()
}
