package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	Id               bson.ObjectId `json:"-" bson:"_id,omitempty"`
	UID              string        `json:"uid,omitempty" bson:"uid,omitempty"` // Firebase UID
	DisplayName      string        `json:"displayName,omitempty" bson:"displayName,omitempty"`
	BirthDate        time.Time     `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	DRE              string        `json:"dre,omitempty" bson:"dre,omitempty"`
	SIAPE            string        `json:"siape,omitempty" bson:"siape,omitempty"`
	Email            string        `json:"email,omitempty" bson:"email,omitempty"`
	Password         string        `json:"password,omitempty" bson:"password,omitempty"`
	CurriculumGridId interface{}   `json:"curriculum,omitempty" bson:"curriculum,omitempty"` // ID do arquivo no grid
	Phone            string        `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Photo            string        `json:"photo,omitempty" bson:"photo,omitempty"`       // URL para foto
	Type             string        `json:"userType,omitempty" bson:"userType,omitempty"` // ALUNO, PROFESSOR, REPRESENTANTE DE EMPRESA/ORGANIZACAO
	GraduationLevel  string        `json:"graduationLevel,omitempty" bson:"graduationLevel,omitempty"`
	Course           string        `json:"course,omitempty" bson:"course,omitempty"`
	Status           string        `json:"status,omitempty" bson:"status,omitempty"` // CURSANDO, CONCLUIDO
	LinkedInProfile  string        `json:"linkedinUrl,omitempty" bson:"linkedinUrl,omitempty"`
	AreaOfInterest   string        `json:"areaOfInterest,omitempty" bson:"areaOfInterest,omitempty"` // TODO Areas of interest?
	CreatedAt        time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt        time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type UserResponse struct {
	UID              string        `json:"uid,omitempty" bson:"uid,omitempty"` // Firebase UID
	DisplayName      string        `json:"displayName,omitempty" bson:"displayName,omitempty"`
	BirthDate        time.Time     `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	DRE              string        `json:"dre,omitempty" bson:"dre,omitempty"`
	SIAPE            string        `json:"siape,omitempty" bson:"siape,omitempty"`
	Email            string        `json:"email,omitempty" bson:"email,omitempty"`
	Phone            string        `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Photo            string        `json:"photo,omitempty" bson:"photo,omitempty"`       // URL para foto
	Type             string        `json:"userType,omitempty" bson:"userType,omitempty"` // ALUNO, PROFESSOR, REPRESENTANTE DE EMPRESA/ORGANIZACAO
	GraduationLevel  string        `json:"graduationLevel,omitempty" bson:"graduationLevel,omitempty"`
	Course           string        `json:"course,omitempty" bson:"course,omitempty"`
	Status           string        `json:"status,omitempty" bson:"status,omitempty"` // CURSANDO, CONCLUIDO
	LinkedInProfile  string        `json:"linkedinUrl,omitempty" bson:"linkedinUrl,omitempty"`
	AreaOfInterest   string        `json:"areaOfInterest,omitempty" bson:"areaOfInterest,omitempty"` // TODO Areas of interest?
	CreatedAt        time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}