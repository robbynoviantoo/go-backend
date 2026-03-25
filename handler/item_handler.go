package handler

import (
	"net/http"

	"go-backend/models"
	"go-backend/repository"
	"go-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var item models.Item
	c.ShouldBindJSON(&item)

	userID, _ := c.Get("user_id")
	item.UserID = userID.(int)

	repository.CreateItem(item)

	c.JSON(200, gin.H{
		"message": "success",
		"data": item,
	})
}

func GetItems(c *gin.Context) {
	userID, _ := c.Get("user_id")

	items, _ := repository.GetItemsByUser(userID.(int))

	c.JSON(http.StatusOK, items)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	repository.DeleteItem(utils.StringToInt(id), userID.(int))

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}