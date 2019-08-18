package controller

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/usecase/service"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
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
	// Binds the request body to a User struct
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := r.Service.CreateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, http.NoBody)
}

func (r *Resource) SendCurriculum(c echo.Context) error {

	authToken, err := getTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Gets the curriculum file
	curriculum, err := getCurriculumFromFormData(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := r.Service.SendCurriculum(curriculum, authToken); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, http.NoBody)
}

func getTokenFromHeader(c echo.Context) (string, error) {
	authToken := c.Request().Header.Get("Token")
	if len(authToken) == 0 {
		return "", errors.New("authorization token empty")
	}

	return authToken, nil
}

func getCurriculumFromFormData(c echo.Context) (*os.File, error) {

	file, err := c.FormFile("curriculum")
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return dst, nil
}
