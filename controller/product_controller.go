package controller

import (
	"crud-go/model"
	"crud-go/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productController struct {
	productUseCase use_case.ProductUseCase
}

func NewProductController(useCase use_case.ProductUseCase) productController {
	return productController{
		productUseCase: useCase,
	}
}

func (p productController) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, product)

}

func (p *productController) GetAllProducts(c *gin.Context) {
	products, err := p.productUseCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		response := model.Response{
			Message: "Invalid Product ID",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		response := model.Response{
			Message: "Invalid Product ID",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product.ID = productIdInt

	product, err = p.productUseCase.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}
