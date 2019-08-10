package controller

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/usecase/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Resource struct {
	Service service.Vagas
}

func NewController(service service.Vagas) Vagas {
	return &Resource{
		Service: service,
	}
}

func (r *Resource) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := r.Service.CreateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, http.NoBody)
}