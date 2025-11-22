package handler

import (
	// "fmt"
	"net/http"

	"github.com/davidcm146/shopee-microservice/product-service/internal/common/errors"
	"github.com/davidcm146/shopee-microservice/product-service/internal/dto"
	"github.com/davidcm146/shopee-microservice/product-service/internal/service"
	"github.com/davidcm146/shopee-microservice/product-service/internal/validation"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func (h *ProductHandler) ProductService() service.ProductService {
	return h.productService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.FindAll(c.Request.Context())
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.FindByID(c.Request.Context(), id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *ProductHandler) GetProductBySellerID(c *gin.Context) {
	currentUser, ok := c.Get("currentUser")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := currentUser.(struct {
		ID    string `json:"id"`
		Role  string `json:"role"`
		Email string `json:"email"`
	})
	products, err := h.productService.FindBySellerID(c.Request.Context(), user.ID)

	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input dto.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if errs, err := validation.ValidateStruct(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	product, err := h.productService.Create(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": product, "message": "Product created successfully"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var input dto.UpdateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	id := c.Param("id")
	if errs, err := validation.ValidateStruct(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	updatedProduct, err := h.productService.Update(c.Request.Context(), id, &input)

	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{"error": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": updatedProduct, "message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.productService.SoftDelete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
