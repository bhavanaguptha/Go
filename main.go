package main

import (
	"go_practice/config"
	routes "go_practice/routers"
	"os"
)

func main() {

	db := config.SetupDB()
	// db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run(":" + os.Getenv("PORT"))
}
