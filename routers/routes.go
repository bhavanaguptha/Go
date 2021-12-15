package routes

import (
	"go_practice/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	usersGroup := r.Group("users")
	{
		usersGroup.POST("register", services.Registration)
		usersGroup.GET("getAllUsers", services.GetUsers)
		usersGroup.GET("getsingleuser/:id", services.GetSingleUser)
		usersGroup.POST("updateuser/:id", services.UpdateUser)
	}
	return r
}
