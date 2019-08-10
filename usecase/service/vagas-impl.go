package service

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/persistance/repository"
	"firebase.google.com/go/auth"
)

type Resource struct {
	Repository repository.Vagas
	auth       *auth.Client
}

func NewService(repository repository.Vagas, auth *auth.Client) Vagas {
	return &Resource{
		Repository: repository,
		auth: auth,
	}
}

func (r *Resource) CreateUser(user *model.User) error {
	return nil
}