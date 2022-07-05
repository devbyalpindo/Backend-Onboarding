package middleware

import (
	"fse-onboarding/helper"
	"fse-onboarding/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func ViewerAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		userId, _, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
		if err != nil {
			resp := helper.ResponseError("You are unathorized", "Invalid token", 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		c.Set("user_id", userId)

	}
}
