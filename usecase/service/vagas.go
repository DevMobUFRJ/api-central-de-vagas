package service

import (
	"api-central-de-vagas/model"
	"mime/multipart"
)

type Vagas interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, authToken string) error
	SendCurriculum(curriculum multipart.File, authToken string) error
	CreateVaga(vaga *model.Vaga, authToken string) error
}
