package controller

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/usecase/service"
	"errors"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"strings"
)

type Resource struct {
	Service service.Vagas
}

type Vagas interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	SendCurriculum(c echo.Context) error
	CreateVaga(c echo.Context) error
	GetUsers(c echo.Context) error
	GetUserById(c echo.Context) error
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

	return c.JSON(http.StatusCreated, http.NoBody)
}

func (r *Resource) GetUsers(c echo.Context) error {
	users, err := r.Service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, http.NoBody)
	}

	if len(*users) == 0 {
		return c.JSON(http.StatusNoContent, http.NoBody)
	}

	return c.JSON(http.StatusOK, users)
}

func (r *Resource) GetUserById(c echo.Context) error {
	uid := c.Param("uid")
	if len(uid) == 0 {
		return c.JSON(http.StatusBadRequest, errors.New("user uid must not be empty"))
	}

	user, err := r.Service.GetUserByUID(uid)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNoContent, http.NoBody)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (r *Resource) UpdateUser(c echo.Context) error {

	authToken, err := getTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := r.Service.UpdateUser(user, authToken); err != nil {
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

func (r *Resource) CreateVaga(c echo.Context) error {

	authToken, err := getTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	vaga := new(model.Vaga)
	if err := c.Bind(vaga); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := r.Service.CreateVaga(vaga, authToken); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, http.NoBody)
}

func getTokenFromHeader(c echo.Context) (string, error) {
	authToken := c.Request().Header.Get("Token")
	if len(authToken) == 0 {
		return "", errors.New("authorization token empty")
	}

	return authToken, nil
}

func getCurriculumFromFormData(c echo.Context) (multipart.File, error) {

	file, err := c.FormFile("curriculum")
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	return src, nil
}
