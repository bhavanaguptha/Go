package services

import (
	"go_practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Registration(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Registered successfully"})
}

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var output []models.Users
	if err := db.Raw("select * from users").Scan(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

func GetSingleUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var output models.Users
	if err := db.Raw("select * from users where id = ?", c.Param("id")).Scan(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user available on this ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": output})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var output models.Users

	if err := db.Raw("select * from users where id = ?", c.Param("id")).Scan(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No User Found"})
		return
	}

	var input models.Name
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&output).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": output})
}
