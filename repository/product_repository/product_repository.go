package product_repository

import "fse-onboarding/model/entity"

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetDetailProduct(string) (*entity.Product, error)
	AddProduct(entity.Product) (*string, error)
	UpdateProduct(entity.Product, string) (*string, error)
	DeleteProduct(string) error
	CheckedProduct(entity.Product, string) (*string, error)
	PublishProduct(entity.Product, string) (*string, error)
}
