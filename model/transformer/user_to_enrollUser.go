package transformer

import (
	"api-central-de-vagas/model"
)

func UserToEnrollUser(user *model.User) *model.EnrollUser {
	return &model.EnrollUser{
		Id: user.Id,
		UID: user.UID,
		DisplayName: user.DisplayName,
		BirthDate: user.BirthDate,
		CurriculumGridId: user.CurriculumGridId,
		Type: user.Type,
		GraduationLevel: user.GraduationLevel,
		Course: user.Course,
		LinkedInProfile: user.LinkedInProfile,
	}
}
