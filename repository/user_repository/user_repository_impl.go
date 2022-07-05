package user_repository

import (
	"fse-onboarding/helper"
	"fse-onboarding/model/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB}
}

func (repository *UserRepositoryImpl) GetAllUsers() ([]entity.User, error) {
	user := []entity.User{}
	err := repository.DB.Model(&entity.User{}).Preload("Role").Find(&user).Error
	helper.PanicIfError(err)
	if len(user) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (repository *UserRepositoryImpl) GetDetailUsers(id string) (*entity.User, error) {
	user := entity.User{}
	result := repository.DB.Where("id = ?", id).Preload("Role").Find(&user)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) AddUsers(user entity.User) (*string, error) {
	if err := repository.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user.Id, nil
}

func (repository *UserRepositoryImpl) UpdateUsers(user entity.User, id string) (string, error) {
	result := repository.DB.Model(&user).Where("id = ?", id).Updates(entity.User{PersonalNumber: user.PersonalNumber, Password: user.Password, Email: user.Email, RoleID: user.RoleID, Active: user.Active})

	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return id, nil
}

func (repository *UserRepositoryImpl) DeleteUsers(id string) error {
	result := repository.DB.Where("id = ?", id).Delete(&entity.User{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repository *UserRepositoryImpl) LoginUsers(pn string) (*entity.User, error) {
	user := entity.User{}
	result := repository.DB.Where("personal_number = ?", pn).Find(&user)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) GetAllRole() ([]entity.Role, error) {
	role := []entity.Role{}
	err := repository.DB.Model(&entity.Role{}).Scan(&role).Error
	helper.PanicIfError(err)
	if len(role) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return role, nil
}

func (repository *UserRepositoryImpl) GetRoleById(id string) (*entity.Role, error) {
	role := entity.Role{}
	result := repository.DB.Where("id = ?", id).Find(&role)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &role, nil
}

func (repository *UserRepositoryImpl) GetRoleIDByName(name string) (*entity.Role, error) {
	role := entity.Role{}
	result := repository.DB.Where("title = ?", name).Find(&role)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &role, nil
}
