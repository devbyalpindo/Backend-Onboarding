package user_repository

import "fse-onboarding/model/entity"

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetDetailUsers(string) (*entity.User, error)
	AddUsers(entity.User) (*string, error)
	UpdateUsers(entity.User, string) (string, error)
	DeleteUsers(string) error
	LoginUsers(personal_number string) (*entity.User, error)
	GetAllRole() ([]entity.Role, error)
	GetRoleById(id string) (*entity.Role, error)
	GetRoleIDByName(name string) (*entity.Role, error)
}
