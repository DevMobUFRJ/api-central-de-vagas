package service

import (
	"api-central-de-vagas/model"
	"mime/multipart"
)

type Vagas interface {
	CreateUser(user *model.User) error
	SendCurriculum(curriculum multipart.File, authToken string) error
}
