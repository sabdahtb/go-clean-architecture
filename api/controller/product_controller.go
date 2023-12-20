package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sabda.go/learn/bootstrap"
	"sabda.go/learn/domain"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
	Env            *bootstrap.Env
	DB             *gorm.DB
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var (
		request domain.CreateProductRequest
	)

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	product := domain.Product{
		Name:  request.Name,
		Stock: request.Stock,
	}
	result, err := pc.ProductUsecase.Create(c, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := &domain.SuccessResponse{
		Message: result,
	}
	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetAllProduct(c *gin.Context) {
	response, err := pc.ProductUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")
	response, err := pc.ProductUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	result, err := pc.ProductUsecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := &domain.SuccessResponse{
		Message: result,
	}

	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var (
		request domain.CreateProductRequest
	)
	id := c.Param("id")

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	product := domain.Product{
		Name:  request.Name,
		Stock: request.Stock,
	}
	result, err := pc.ProductUsecase.Update(c, id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := &domain.SuccessResponse{
		Message: result,
	}
	c.JSON(http.StatusOK, response)
}
