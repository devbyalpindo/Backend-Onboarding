package product_repository

import (
	"fse-onboarding/helper"
	"fse-onboarding/model/entity"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB}
}

func (repository *ProductRepositoryImpl) GetAllProducts() ([]entity.Product, error) {
	product := []entity.Product{}
	err := repository.DB.Model(&entity.Product{}).Scan(&product).Error
	helper.PanicIfError(err)
	if len(product) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) GetDetailProduct(id string) (*entity.Product, error) {
	product := entity.Product{}
	result := repository.DB.Where("id = ?", id).Preload("Maker").Preload("Checker").Preload("Signer").Find(&product)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &product, nil
}

func (repository *ProductRepositoryImpl) AddProduct(product entity.Product) (*string, error) {
	if err := repository.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product.Id, nil
}

func (repository *ProductRepositoryImpl) UpdateProduct(product entity.Product, id string) (*string, error) {
	result := repository.DB.Model(&product).Where("id = ?", id).Updates(entity.Product{Name: product.Name, Description: product.Description})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &id, nil
}

func (repository *ProductRepositoryImpl) DeleteProduct(id string) error {
	result := repository.DB.Where("id = ?", id).Delete(&entity.Product{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repository *ProductRepositoryImpl) CheckedProduct(product entity.Product, id string) (*string, error) {
	result := repository.DB.Model(&product).Where("id = ?", id).Updates(entity.Product{Name: product.Name, Description: product.Description, Status: "approved", CheckerID: product.CheckerID})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &id, nil
}

func (repository *ProductRepositoryImpl) PublishProduct(product entity.Product, id string) (*string, error) {
	result := repository.DB.Model(&product).Where("id = ?", id).Updates(entity.Product{Name: product.Name, Description: product.Description, Status: "active", SignerID: product.SignerID})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &id, nil
}
