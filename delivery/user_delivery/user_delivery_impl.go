package user_delivery

import (
	"fse-onboarding/helper"
	"fse-onboarding/model/dto"
	"fse-onboarding/usecase/user_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDeliveryImpl struct {
	usecase user_usecase.UserUsecase
}

func NewUserDelivery(userUsecase user_usecase.UserUsecase) UserDelivery {
	return &UserDeliveryImpl{userUsecase}
}

func (res *UserDeliveryImpl) GetAllUsers(c *gin.Context) {

	response := res.usecase.GetAllUsers()
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *UserDeliveryImpl) GetDetailUsers(c *gin.Context) {
	userID := c.Param("id")

	response := res.usecase.GetDetailUsers(userID)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *UserDeliveryImpl) AddUsers(c *gin.Context) {
	userReq := dto.UserRequest{}
	if err := c.ShouldBindJSON(&userReq); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.AddUsers(userReq)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (res *UserDeliveryImpl) UpdateUsers(c *gin.Context) {
	id := c.Param("id")

	userRequest := dto.UserRequest{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.UpdateUsers(userRequest, id)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *UserDeliveryImpl) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUsers(id)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (res *UserDeliveryImpl) UserLogin(c *gin.Context) {
	userLogin := dto.UserLogin{}
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.LoginUsers(userLogin)

	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (res *UserDeliveryImpl) GetAllRoles(c *gin.Context) {

	response := res.usecase.GetAllRole()
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
