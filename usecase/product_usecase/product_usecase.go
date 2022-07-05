package product_usecase

import (
	"fse-onboarding/model/dto"
)

type ProductUsecase interface {
	GetAllProducts() dto.Response
	GetDetailProduct(id string) dto.Response
	AddProduct(dto.ProductRequest, string) dto.Response
	UpdateProduct(id string, payload dto.ProductRequest) dto.Response
	DeleteProduct(id string) dto.Response
	CheckedProduct(payload dto.ProductRequest, id string, userID string) dto.Response
	PublishProduct(payload dto.ProductRequest, id string, userID string) dto.Response
}
