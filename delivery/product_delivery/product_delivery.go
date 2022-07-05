package product_delivery

import "github.com/gin-gonic/gin"

type ProductDelivery interface {
	GetAllProducts(*gin.Context)
	GetDetailProduct(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	CheckedProduct(*gin.Context)
	PublishProduct(*gin.Context)
}
