package router

import (
	"fse-onboarding/delivery/product_delivery"
	"fse-onboarding/delivery/user_delivery"
	"fse-onboarding/middleware"
	"fse-onboarding/usecase/jwt_usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(jwtUsecase jwt_usecase.JwtUsecase, userDelivery user_delivery.UserDelivery, productDelivery product_delivery.ProductDelivery) *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.POST("/login", userDelivery.UserLogin)

	//viewer middleware
	viewerRoute := router.Group("/")
	viewerRoute.Use(middleware.ViewerAuth(jwtUsecase))
	{
		viewerRoute.GET("/users", userDelivery.GetAllUsers)
		viewerRoute.GET("/users/:id", userDelivery.GetDetailUsers)
		viewerRoute.GET("/roles", userDelivery.GetAllRoles)
		viewerRoute.GET("/products", productDelivery.GetAllProducts)
		viewerRoute.GET("/products/:id", productDelivery.GetDetailProduct)
	}

	//maker middleware
	makerRoute := router.Group("/")
	makerRoute.Use(middleware.MakerAuth(jwtUsecase))
	{
		makerRoute.POST("/products", productDelivery.AddProduct)

	}

	//checker middleware
	checkerRoute := router.Group("/")
	checkerRoute.Use(middleware.CheckerAuth(jwtUsecase))
	{
		checkerRoute.PUT("/products/:id/checked", productDelivery.CheckedProduct)
	}

	//signer middleware
	signerRoute := router.Group("/")
	signerRoute.Use(middleware.SignerAuth(jwtUsecase))
	{
		signerRoute.PUT("/products/:id/published", productDelivery.PublishProduct)
	}

	//admin middleware
	adminRoute := router.Group("/")
	adminRoute.Use(middleware.AdminAuth(jwtUsecase))
	{
		adminRoute.PUT("/users/:id", userDelivery.UpdateUsers)
		adminRoute.DELETE("/users/:id", userDelivery.DeleteUsers)
		adminRoute.POST("/users", userDelivery.AddUsers)

		adminRoute.PUT("/products/:id", productDelivery.UpdateProduct)
		adminRoute.DELETE("/products/:id", productDelivery.DeleteProduct)
	}

	return router
}
