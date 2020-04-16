package repository

import (
	"api-central-de-vagas/model"
	"firebase.google.com/go/storage"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

type Resource struct {
	mongoSession *mgo.Session
	storage *storage.Client
}

type Vagas interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	SendCurriculum(curriculum multipart.File, userName string) (interface{}, error)
	GetCurriculumByGridId(gridId interface{}) ([]byte, error)
	CreateVaga(vaga *model.Vaga) error
	GetUserByUID(uid string) (*model.User, error)
	GetUserResponseByUID(uid string) (*model.UserResponse, error)
	GetUsers() (*[]model.UserResponse, error)
}

func NewRepository(mongoSession *mgo.Session, storage *storage.Client) Vagas {
	return &Resource{
		mongoSession: mongoSession,
		storage: storage,
	}
}

func (r *Resource) GetUsers() (*[]model.UserResponse, error) {
	users := new([]model.UserResponse)

	if err := r.getUsersCollection().Find(nil).All(users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Resource) CreateUser(user *model.User) error {
	if err := r.getUsersCollection().Insert(user); err != nil {
		return err
	}

	return nil
}

func (r *Resource) CreateVaga(vaga *model.Vaga) error {
	if err := r.getVagasCollection().Insert(vaga); err != nil {
		return err
	}

	return nil
}

func (r *Resource) UpdateUser(user *model.User) error {
	set := bson.M{}
	nilTime := time.Time{}

	if user.DisplayName != "" {
		set["displayName"] = user.DisplayName
	}
	if user.BirthDate != nilTime {
		set["birthdate"] = user.BirthDate
	}
	if user.DRE != "" {
		set["dre"] = user.DRE
	}
	if user.SIAPE != "" {
		set["siape"] = user.SIAPE
	}
	if user.Email != "" {
		set["email"] = user.Email
	}
	if user.Password != "" {
		set["password"] = user.Password
	}
	if user.CurriculumGridId != nil {
		set["curriculum"] = user.CurriculumGridId
	}
	if user.Phone != "" {
		set["phoneNumber"] = user.Phone
	}
	if user.Photo != "" {
		set["photo"] = user.Photo
	}
	if user.Type != "" {
		set["userType"] = user.Type
	}
	if user.GraduationLevel != "" {
		set["graduationLevel"] = user.GraduationLevel
	}
	if user.Course != "" {
		set["course"] = user.Course
	}
	if user.Status != "" {
		set["status"] = user.Status
	}
	if user.LinkedInProfile != "" {
		set["linkedinUrl"] = user.LinkedInProfile
	}
	if user.AreaOfInterest != "" {
		set["areaOfInterest"] = user.AreaOfInterest
	}

	update := bson.M{"$set": set}
	query := bson.M{"uid": user.UID}

	if err := r.getUsersCollection().Update(query, update); err != nil {
		return err
	}

	return nil
}

func (r *Resource) SendCurriculum(curriculum multipart.File, userName string) (interface{}, error) {
	defer curriculum.Close()

	file, err := r.mongoSession.DB("central-de-vagas").GridFS("fs").Create(fmt.Sprintf("CV_%s.pdf", userName))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, curriculum)
	if err != nil {
		return "", err
	}

	if err := file.Close(); err != nil {
		return "", err
	}

	return file.Id(), nil

	/*Firebase
	storageTime := time.Now()
	bucket, err := r.storage.Bucket("central-de-vagas.appspot.com")
	wc := bucket.Object(fmt.Sprintf("CV_%s.pdf", userName)).NewWriter(context.Background())
	_, err = io.Copy(wc, curriculum)
	if err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}
	fmt.Println("Time for storage: ", time.Since(storageTime))
	*/
}

func (r *Resource) GetCurriculumByGridId(gridId interface{}) ([]byte, error) {
	query := bson.M{
		"files_id": bson.ObjectIdHex(strings.Split(fmt.Sprintf("%v", gridId), `"`)[1]),
	}

	file := &struct {
		Data []byte `bson:"data"`
	}{}

	if err := r.mongoSession.DB("central-de-vagas").GridFS("fs").Chunks.Find(query).One(file); err != nil {
		return nil, err
	}

	return file.Data, nil
}

func (r *Resource) GetUserByUID(uid string) (*model.User, error) {
	query := bson.M{
		"uid": uid,
	}

	user := new(model.User)

	if err := r.getUsersCollection().Find(query).One(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resource) GetUserResponseByUID(uid string) (*model.UserResponse, error) {
	query := bson.M{
		"uid": uid,
	}

	user := new(model.UserResponse)

	if err := r.getUsersCollection().Find(query).One(user); err != nil {
		return nil, err
	}

	return user, nil
}

/*
func (r *Resource) GetCurriculum(photo *model.Photo) error {

}
*/

func (r *Resource) getUsersCollection() *mgo.Collection {
	return r.mongoSession.DB("central-de-vagas").C("users")
}

func (r *Resource) getVagasCollection() *mgo.Collection {
	return r.mongoSession.DB("central-de-vagas").C("vagas")
}