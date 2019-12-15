package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Vaga struct {
	Title       string       `json:"title,omitempty" bson:"title,omitempty"`
	Description string       `json:"description,omitempty" bson:"description,omitempty"`
	Type        string       `json:"type,omitempty" bson:"type,omitempty"`
	Hours       int64        `json:"hours,omitempty" bson:"hours,omitempty"`
	Benefits    []string     `json:"benefits,omitempty" bson:"benefits,omitempty"`
	Income      float64      `json:"income,omitempty" bson:"income,omitempty"`
	Enrolled    []EnrollUser `json:"enrolled,omitempty" bson:"enrolled,omitempty"`
	Location    string       `json:"location,omitempty" bson:"location,omitempty"`
	Creator		string 		 `json:"creatorUID,omitempty" bson:"creatorUID,omitempty"` // UID do criador da vaga
	Active      bool		 `json:"active" bson:"active"`
}

type EnrollUser struct {
	Id               bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	UID              string        `json:"uid,omitempty" bson:"uid,omitempty"` // Firebase UID
	DisplayName      string        `json:"displayName,omitempty" bson:"displayName,omitempty"`
	BirthDate        time.Time     `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	CurriculumGridId interface{}   `json:"curriculum,omitempty" bson:"curriculum,omitempty"` // ID do arquivo no grid
	Type             string        `json:"userType,omitempty" bson:"userType,omitempty"`     // ALUNO, PROFESSOR, REPRESENTANTE DE EMPRESA/ORGANIZACAO
	GraduationLevel  string        `json:"graduationLevel,omitempty" bson:"graduationLevel,omitempty"`
	Course           string        `json:"course,omitempty" bson:"course,omitempty"`
	LinkedInProfile  string        `json:"linkedinUrl,omitempty" bson:"linkedinUrl,omitempty"`
}
