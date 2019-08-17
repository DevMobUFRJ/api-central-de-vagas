package service

import (
	"api-central-de-vagas/model"
	"os"
)

type Vagas interface {
	CreateUser(user *model.User) error
	SendCurriculum(curriculum *os.File, authToken string) error
}
