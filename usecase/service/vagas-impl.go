package service

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/persistance/repository"
	"context"
	"firebase.google.com/go/auth"
)

type Resource struct {
	Repository repository.Vagas
	Client     *auth.Client
}

func NewService(repository repository.Vagas, client *auth.Client) Vagas {
	return &Resource{
		Repository: repository,
		Client: client,
	}
}

func (r *Resource) CreateUser(user *model.User) error {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		PhoneNumber(user.Phone).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.Photo).
		Disabled(false)

	// Creates the user in firebase
	_, err := r.Client.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	// token, _ := r.Client.VerifyIDToken()

	user.Password = ""
	if err := r.Repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (r *Resource) CreateVaga(vaga *model.Vaga, idToken string) error {
	if err := r.VerifyIDToken(idToken); err != nil {
		return err // Falha na autenticação
	}

	if err := r.Repository.CreateVaga(vaga); err != nil {
		return err
	}

	return nil
}

func (r *Resource) VerifyIDToken(idToken string) error {
	_, err := r.Client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return err
	}

	return nil
}