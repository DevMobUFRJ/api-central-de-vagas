package controller

import "github.com/labstack/echo/v4"

type Vagas interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	SendCurriculum(c echo.Context) error
	CreateVaga(c echo.Context) error
}
