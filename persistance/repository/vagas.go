package repository

import (
	"api-central-de-vagas/model"
	"mime/multipart"
)

type Vagas interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	SendCurriculum(curriculum multipart.File, userName string) (interface{}, error)
	GetUserByUID(uid string) (*model.User, error)
}
