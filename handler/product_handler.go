package handler

import (
	"database/sql"
	"net/http"

	"go-backend/models"
	"go-backend/repository"
	"go-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateProductCategory(c *gin.Context) {
	var category models.ProductCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateProductCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

func GetProductCategories(c *gin.Context) {
	categories, err := repository.GetProductCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func UpdateProductCategory(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))

	var category models.ProductCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateProductCategory(id, category); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated"})
}

func DeleteProductCategory(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))

	if err := repository.DeleteProductCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	product.UserID = userID.(int)
	id, err := repository.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = int(id)
	c.JSON(http.StatusCreated, gin.H{
		"message": "product created",
		"data":    product,
	})
}

func GetProducts(c *gin.Context) {
	products, err := repository.GetProducts(
		c.Query("name"),
		c.Query("category_id"),
		c.Query("user_id"),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductByID(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))

	product, err := repository.GetProductByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := repository.UpdateProduct(id, userID.(int), product); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func DeleteProduct(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := repository.DeleteProduct(id, userID.(int)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
