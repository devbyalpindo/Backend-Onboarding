package product_usecase

import (
	"errors"
	"fse-onboarding/helper"
	"fse-onboarding/model/dto"
	"fse-onboarding/model/entity"
	"fse-onboarding/repository/product_repository"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductUsecaseImpl struct {
	ProductRepository product_repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductUsecase(productRepository product_repository.ProductRepository, validate *validator.Validate) ProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (usecase *ProductUsecaseImpl) GetAllProducts() dto.Response {
	productList, err := usecase.ProductRepository.GetAllProducts()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err, 404)
	} else if err != nil {
		return helper.ResponseError("failed", err, 500)
	}
	helper.PanicIfError(err)
	response := []dto.Product{}
	for _, user := range productList {
		responseData := dto.Product{
			Id:          user.Id,
			Name:        user.Name,
			Description: user.Description,
			Status:      user.Status,
		}
		response = append(response, responseData)
	}

	return helper.ResponseSuccess("ok", nil, response, 200)
}

func (usecase *ProductUsecaseImpl) GetDetailProduct(id string) dto.Response {
	product, err := usecase.ProductRepository.GetDetailProduct(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", "Data not found", 404)
	} else if err != nil {
		return helper.ResponseError("failed", err, 500)
	}
	helper.PanicIfError(err)

	responseData := dto.ProductDetail{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Status:      product.Status,
		Maker: dto.UserAction{
			Id:   product.Maker.Id,
			Name: product.Maker.Name,
		},
		Checker: dto.UserAction{
			Id:   product.Checker.Id,
			Name: product.Checker.Name,
		},
		Signer: dto.UserAction{
			Id:   product.Signer.Id,
			Name: product.Signer.Name,
		},
	}

	return helper.ResponseSuccess("ok", nil, responseData, 200)
}

func (usecase *ProductUsecaseImpl) AddProduct(body dto.ProductRequest, userID string) dto.Response {
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

	payloadProduct := entity.Product{
		Id:          createID,
		Name:        body.Name,
		Description: body.Description,
		MakerID:     userID,
	}

	product, err := usecase.ProductRepository.AddProduct(payloadProduct)
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"id": product}, 201)

}

func (usecase *ProductUsecaseImpl) UpdateProduct(id string, body dto.ProductRequest) dto.Response {
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

	payloadProduct := entity.Product{
		Name:        body.Name,
		Description: body.Description,
	}

	productID, err := usecase.ProductRepository.UpdateProduct(payloadProduct, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	}
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]any{"id": productID}, 200)

}

func (usecase *ProductUsecaseImpl) DeleteProduct(id string) dto.Response {
	err := usecase.ProductRepository.DeleteProduct(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	} else if err != nil {
		return helper.ResponseError("failed", err.Error(), 500)
	}
	return helper.ResponseSuccess("ok", nil, nil, 200)

}

func (usecase *ProductUsecaseImpl) CheckedProduct(body dto.ProductRequest, id string, userID string) dto.Response {
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

	payloadProduct := entity.Product{
		Name:        body.Name,
		Description: body.Description,
		CheckerID:   &userID,
	}

	productID, err := usecase.ProductRepository.CheckedProduct(payloadProduct, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	}
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]any{"id": productID}, 200)
}

func (usecase *ProductUsecaseImpl) PublishProduct(body dto.ProductRequest, id string, userID string) dto.Response {
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

	payloadProduct := entity.Product{
		Name:        body.Name,
		Description: body.Description,
		SignerID:    &userID,
	}

	productID, err := usecase.ProductRepository.PublishProduct(payloadProduct, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("failed", err.Error(), 404)
	}
	helper.PanicIfError(err)

	return helper.ResponseSuccess("ok", nil, map[string]any{"id": productID}, 200)

}
