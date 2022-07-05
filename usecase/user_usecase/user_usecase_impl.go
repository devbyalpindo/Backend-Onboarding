package user_usecase

import (
	"errors"
	"fse-onboarding/helper"
	"fse-onboarding/model/dto"
	"fse-onboarding/model/entity"
	"fse-onboarding/repository/user_repository"
	"fse-onboarding/usecase/jwt_usecase"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserUsecaseImpl struct {
	UserRepository user_repository.UserRepository
	JwtUsecase     jwt_usecase.JwtUsecase
	Validate       *validator.Validate
}

func NewUserUsecase(userRepository user_repository.UserRepository, jwtUsecase jwt_usecase.JwtUsecase, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
		JwtUsecase:     jwtUsecase,
		Validate:       validate,
	}
}

func (usecase *UserUsecaseImpl) GetAllUsers() dto.Response {
	userList, err := usecase.UserRepository.GetAllUsers()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	helper.PanicIfError(err)
	response := []dto.User{}
	for _, user := range userList {
		responseData := dto.User{
			Id:             user.Id,
			PersonalNumber: user.PersonalNumber,
			Email:          user.Email,
			Name:           user.Name,
			Role: dto.Role{
				Id:    user.Role.Id,
				Title: user.Role.Title,
			},
			Active: user.Active,
		}
		response = append(response, responseData)
	}

	return helper.ResponseSuccess("ok", nil, response, 200)
}

func (usecase *UserUsecaseImpl) GetDetailUsers(id string) dto.Response {
	user, err := usecase.UserRepository.GetDetailUsers(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", "Data not found", 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	helper.PanicIfError(err)

	responseData := dto.User{
		Id:             user.Id,
		PersonalNumber: user.PersonalNumber,
		Email:          user.Email,
		Name:           user.Name,
		Role: dto.Role{
			Id:    user.Role.Id,
			Title: user.Role.Title,
		},
		Active: user.Active,
	}

	return helper.ResponseSuccess("ok", nil, responseData, 200)
}

func (usecase *UserUsecaseImpl) AddUsers(body dto.UserRequest) dto.Response {
	err := usecase.Validate.Struct(body)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]dto.CustomMessageError, len(ve))
			for i, fe := range ve {
				out[i] = dto.CustomMessageError{
					Field:    fe.Field(),
					Messsage: helper.MessageError(fe.Tag()),
				}
			}
			return helper.ResponseError("failed", out, 403)
		}

	}

	createID := uuid.New().String()
	encryptPwd, err := helper.HashPassword(body.Password)
	helper.PanicIfError(err)

	role, err := usecase.UserRepository.GetRoleIDByName("viewer")
	helper.PanicIfError(err)

	payloadUser := entity.User{
		Id:             createID,
		PersonalNumber: body.PersonalNumber,
		Password:       encryptPwd,
		Name:           body.Name,
		Email:          body.Email,
		RoleID:         role.Id,
	}

	user, err := usecase.UserRepository.AddUsers(payloadUser)
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"id": user}, 201)

}

func (usecase *UserUsecaseImpl) UpdateUsers(body dto.UserRequest, id string) dto.Response {
	err := usecase.Validate.Struct(body)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]dto.CustomMessageError, len(ve))
			for i, fe := range ve {
				out[i] = dto.CustomMessageError{
					Field:    fe.Field(),
					Messsage: helper.MessageError(fe.Tag()),
				}
			}
			return helper.ResponseError("failed", out, 403)
		}

	}
	encryptPwd, err := helper.HashPassword(body.Password)
	helper.PanicIfError(err)

	payloadUser := entity.User{
		PersonalNumber: body.PersonalNumber,
		Password:       encryptPwd,
		Email:          body.Email,
		Name:           body.Name,
		RoleID:         body.RoleID,
	}

	user, err := usecase.UserRepository.UpdateUsers(payloadUser, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	}
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]any{"id": user}, 200)

}

func (usecase *UserUsecaseImpl) DeleteUsers(id string) dto.Response {
	err := usecase.UserRepository.DeleteUsers(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	} else if err != nil {
		return helper.ResponseError("failed", err.Error(), 500)
	}
	return helper.ResponseSuccess("ok", nil, nil, 200)

}

func (usecase *UserUsecaseImpl) LoginUsers(userPayload dto.UserLogin) dto.Response {
	user, err := usecase.UserRepository.LoginUsers(userPayload.PersonalNumber)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", "Username or password incorrect", 404)
	}

	errPassword := helper.CheckPasswordHash(userPayload.Password, user.Password)

	if errPassword != nil {
		return helper.ResponseError("failed", "Username or password incorrect", 400)
	}

	if !user.Active {
		return helper.ResponseError("failed", "Your account is not active", 400)
	}

	jwt, err := usecase.JwtUsecase.GenerateToken(user.Id, user.RoleID)

	if err != nil {
		return helper.ResponseError("failed", "Wrong personal number / password", 404)
	}

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"token": jwt}, 200)

}

func (usecase *UserUsecaseImpl) GetAllRole() dto.Response {
	roleList, err := usecase.UserRepository.GetAllRole()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", "Data not found", 404)
	} else if err != nil {
		return helper.ResponseError("failed", err, 500)
	}
	helper.PanicIfError(err)
	response := []dto.Role{}
	for _, role := range roleList {
		responseData := dto.Role{
			Id:    role.Id,
			Title: role.Title,
		}
		response = append(response, responseData)
	}

	return helper.ResponseSuccess("ok", nil, response, 200)
}
