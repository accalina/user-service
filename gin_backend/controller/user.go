package controller

import (
	"gin_backend/model"
	"gin_backend/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var users []model.User
	model.DB.Where("is_delete = ?", false).Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users, "success": true})
}

func CreateUser(c *gin.Context) {
	// Validate userinput
	var serializer serializer.UserCreateSerializer
	if err := c.ShouldBindJSON(&serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	// Create user
	user := model.User{Username: serializer.Username, Fullname: serializer.Fullname}
	model.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "data created", "success": true})
}

func RetriveUser(c *gin.Context) {
	var user model.User
	if err := model.DB.Where("id = ? AND is_delete = ?", c.Param("id"), false).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "success": true})
}

func UpdateUser(c *gin.Context) {
	// Retrive user data
	var user model.User
	if err := model.DB.Where("id = ? AND is_delete = ?", c.Param("id"), false).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found", "success": false})
		return
	}

	// Validate userinput
	var serializer serializer.UserUpdateSerializer
	if err := c.ShouldBindJSON(&serializer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	userdata := model.User{Username: serializer.Username, Fullname: serializer.Fullname}
	model.DB.Model(&user).Updates(userdata)

	c.JSON(http.StatusOK, gin.H{"message": "data updated", "success": true})
}

func DeleteUser(c *gin.Context) {
	// Retrive user data
	var user model.User
	if err := model.DB.Where("id = ? AND is_delete = ?", c.Param("id"), false).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found", "success": false})
		return
	}

	userdata := model.User{IsDelete: true}
	model.DB.Model(&user).Updates(userdata)

	c.JSON(http.StatusOK, gin.H{"message": "data deleted", "success": true})
}
