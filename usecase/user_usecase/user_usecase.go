package user_usecase

import (
	"fse-onboarding/model/dto"
)

type UserUsecase interface {
	GetAllUsers() dto.Response
	GetDetailUsers(id string) dto.Response
	AddUsers(dto.UserRequest) dto.Response
	UpdateUsers(dto.UserRequest, string) dto.Response
	DeleteUsers(string) dto.Response
	LoginUsers(dto.UserLogin) dto.Response
	GetAllRole() dto.Response
}
