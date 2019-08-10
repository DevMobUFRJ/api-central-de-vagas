package repository

import "api-central-de-vagas/model"

type Vagas interface {
	CreateUser(user *model.User) error
	GetUserByUID(uid string) (*model.User, error)
}
