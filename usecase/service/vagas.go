package service

import (
	"api-central-de-vagas/model"
	"api-central-de-vagas/persistance/repository"
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"mime/multipart"
	"time"
)

type Resource struct {
	Repository repository.Vagas
	Client     *auth.Client
}

type Vagas interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, authToken string) error
	SendCurriculum(curriculum multipart.File, authToken string) error
	CreateVaga(vaga *model.Vaga, authToken string) error
	GetUsers() (*[]model.UserResponse, error)
	GetUserByUID(id string) (*model.UserResponse, error)
}

func NewService(repository repository.Vagas, client *auth.Client) Vagas {
	return &Resource{
		Repository: repository,
		Client: client,
	}
}


func (r *Resource) GetUsers() (*[]model.UserResponse, error) {
	users, err := r.Repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Resource) GetUserByUID(id string) (*model.UserResponse, error) {
	user, err := r.Repository.GetUserResponseByUID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
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
	userRecord, err := r.Client.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	user.UID = userRecord.UID
	user.Password = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := r.Repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (r *Resource) UpdateUser(user *model.User, authToken string) error {
	uuid, err := r.VerifyIDToken(authToken)
	if err != nil {
		return err
	}

	user.UID = uuid

	return r.Repository.UpdateUser(user)
}

func (r *Resource) SendCurriculum(curriculum multipart.File, authToken string) error {
	uuid, err := r.VerifyIDToken(authToken)
	if err != nil {
		return err
	}

	user, err := r.Repository.GetUserByUID(uuid)
	if err != nil {
		return err
	}

	fileId, err := r.Repository.SendCurriculum(curriculum, user.DisplayName)
	if err != nil {
		fmt.Println("Erro send")
		return err
	}

	user.CurriculumGridId = fileId
	user.UpdatedAt = time.Now()

	if err := r.Repository.UpdateUser(user); err != nil {
		fmt.Println("Erro update")
		return err
	}

	return nil
}

func (r *Resource) CreateVaga(vaga *model.Vaga, authToken string) error {
	uuid, err := r.VerifyIDToken(authToken)
	if err != nil {
		return err
	}

	vaga.Creator = uuid
	if err := r.Repository.CreateVaga(vaga); err != nil {
		return err
	}

	return nil
}

func (r *Resource) VerifyIDToken(idToken string) (string, error) {
	token, err := r.Client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}
	
	fmt.Println("User authorized with UID ",  token.UID)

	return token.UID, nil
}