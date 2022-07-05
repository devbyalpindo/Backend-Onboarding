package product_delivery

import (
	"fse-onboarding/helper"
	"fse-onboarding/model/dto"
	"fse-onboarding/usecase/product_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductDeliveryImpl struct {
	usecase product_usecase.ProductUsecase
}

func NewProductDelivery(productUsecase product_usecase.ProductUsecase) ProductDelivery {
	return &ProductDeliveryImpl{productUsecase}
}

func (res *ProductDeliveryImpl) GetAllProducts(c *gin.Context) {

	response := res.usecase.GetAllProducts()
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *ProductDeliveryImpl) GetDetailProduct(c *gin.Context) {
	userID := c.Param("id")

	response := res.usecase.GetDetailProduct(userID)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ProductDeliveryImpl) AddProduct(c *gin.Context) {
	productReq := dto.ProductRequest{}
	userID, _ := c.Get("user_id")
	if err := c.ShouldBindJSON(&productReq); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.AddProduct(productReq, userID.(string))
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (res *ProductDeliveryImpl) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	productRequest := dto.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.UpdateProduct(id, productRequest)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ProductDeliveryImpl) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteProduct(id)
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (res *ProductDeliveryImpl) CheckedProduct(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	productRequest := dto.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.CheckedProduct(productRequest, id, userID.(string))
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ProductDeliveryImpl) PublishProduct(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	productRequest := dto.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errorRes := helper.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.PublishProduct(productRequest, id, userID.(string))
	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
