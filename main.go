package main

import (
	"dessert-shop/database"
	//"dessert-shop/models"
	"dessert-shop/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	//database.DB.AutoMigrate(&models.Dessert{})

	r := gin.Default()

	r.GET("/desserts", handlers.GetDesserts)
	r.POST("/desserts", handlers.CreateDessert)
	r.PUT("/desserts/:id", handlers.UpdateDessert)
	r.DELETE("/desserts/:id", handlers.DeleteDessert)

	r.Run(":8080")
}
