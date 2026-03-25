package handler

import (
	"net/http"

	"go-backend/models"
	"go-backend/repository"
	"go-backend/service"
	"go-backend/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetUsers(c *gin.Context) {
	users, _ := repository.GetAllUsers()
	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	c.ShouldBindJSON(&user)

	repository.UpdateUser(utils.StringToInt(id), user)

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	repository.DeleteUser(utils.StringToInt(id))

	c.JSON(200, gin.H{"message": "deleted"})
}

func Me(c *gin.Context) {
	userID, e := c.Get("user_id")
	if !e {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := repository.Me(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
