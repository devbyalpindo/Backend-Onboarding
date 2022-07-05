package main

import (
	"fse-onboarding/config"
	"fse-onboarding/delivery/product_delivery"
	"fse-onboarding/delivery/user_delivery"
	"fse-onboarding/repository/product_repository"
	"fse-onboarding/repository/user_repository"
	"fse-onboarding/router"
	"fse-onboarding/usecase/jwt_usecase"
	"fse-onboarding/usecase/product_usecase"
	"fse-onboarding/usecase/user_usecase"

	"github.com/go-playground/validator"
)

func main() {
	connection := config.Connect()
	validate := validator.New()

	userRepository := user_repository.NewUserRepository(connection)
	jwtUsecase := jwt_usecase.NewJwtUsecase(userRepository)
	userUsecase := user_usecase.NewUserUsecase(userRepository, jwtUsecase, validate)
	userDelivery := user_delivery.NewUserDelivery(userUsecase)

	productRepository := product_repository.NewProductRepository(connection)
	productUsecase := product_usecase.NewProductUsecase(productRepository, validate)
	productDelivery := product_delivery.NewProductDelivery(productUsecase)

	router := router.NewRouter(jwtUsecase, userDelivery, productDelivery)
	router.Run(":8080")
}
