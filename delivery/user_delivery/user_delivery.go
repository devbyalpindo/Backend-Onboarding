package user_delivery

import "github.com/gin-gonic/gin"

type UserDelivery interface {
	GetAllUsers(*gin.Context)
	GetDetailUsers(*gin.Context)
	AddUsers(*gin.Context)
	UpdateUsers(*gin.Context)
	DeleteUsers(*gin.Context)
	UserLogin(*gin.Context)
	GetAllRoles(*gin.Context)
}
