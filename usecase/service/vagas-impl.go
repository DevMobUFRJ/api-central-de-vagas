package service

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/persistance/repository"
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"os"
	"time"
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

	user.Password = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := r.Repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (r *Resource) SendCurriculum(curriculum *os.File, authToken string) error {
	uuid, err := r.VerifyIDToken(authToken)
	if err != nil {
		return err
	}

	user, err := r.Repository.GetUserByUID(uuid)
	if err != nil {
		return err
	}

	fileId, err := r.Repository.SendCurriculum(curriculum, user.CreatedAt)
	if err != nil {
		return err
	}

	user.CurriculumGridId = fileId
	user.UpdatedAt = time.Now()

	if err := r.Repository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

/*
func (r *Resource) CreateVaga(vaga *model.Vaga, idToken string) error {
	if err := r.VerifyIDToken(idToken); err != nil {
		return err // Falha na autenticação
	}

	if err := r.Repository.CreateVaga(vaga); err != nil {
		return err
	}

	return nil
}
*/

func (r *Resource) VerifyIDToken(idToken string) (string, error) {
	token, err := r.Client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}

	fmt.Println("User authorized with UID ",  token.UID)

	return token.UID, nil
}