package service

import "api-central-de-vagas/model"

type Vagas interface {
	CreateUser(user *model.User) error
}
