package model

import "github.com/globalsign/mgo/bson"

type User struct {
	Id              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	UID             string        `json:"uid,omitempty" bson:"uid,omitempty"` // Firebase UID
	DisplayName     string        `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Email           string        `json:"email,omitempty" bson:"email,omitempty"`
	Phone           string        `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Password        string        `json:"password,omitempty" bson:"password,omitempty"`
	Photo           string        `json:"photo,omitempty" bson:"photo,omitempty"`       // URL para foto
	Type            string        `json:"userType,omitempty" bson:"userType,omitempty"` // ALUNO, PROFESSOR, REPRESENTANTE DE EMPRESA/ORGANIZACAO
	GraduationLevel string        `json:"graduationLevel,omitempty" bson:"graduationLevel,omitempty"`
	Status          string        `json:"status,omitempty" bson:"status,omitempty"` // CURSANDO, CONCLUIDO
}
