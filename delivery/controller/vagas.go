package controller

import "github.com/labstack/echo/v4"

type Vagas interface {
	CreateUser(c echo.Context) error
}
