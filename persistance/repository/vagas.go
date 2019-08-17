package repository

import (
	"api-central-de-vagas/model"
	"os"
	"time"
)

type Vagas interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	SendCurriculum(curriculum *os.File, createdAt time.Time) (interface{}, error)
	GetUserByUID(uid string) (*model.User, error)
}
