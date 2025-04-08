package handlers

import (
	"dessert-shop/database"
	"dessert-shop/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetDesserts(c *gin.Context) {
	var desserts []models.Dessert
	database.DB.Find(&desserts)
	c.JSON(http.StatusOK, desserts)
}

func CreateDessert(c *gin.Context) {
	var dessert models.Dessert
	if err := c.ShouldBindJSON(&dessert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&dessert)
	c.JSON(http.StatusOK, dessert)
}

func UpdateDessert(c *gin.Context) {
	id := c.Param("id")
	var dessert models.Dessert

	if err := database.DB.First(&dessert, id).Error; err != nil {
		log.Panicln("错误", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到该甜品"})
		return
	}

	var input models.Dessert
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dessert.Name = input.Name
	dessert.Description = input.Description
	dessert.Price = input.Price

	if err := database.DB.Save(&dessert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, dessert)
}

func DeleteDessert(c *gin.Context) {
	id := c.Param("id")
	var dessert models.Dessert

	if err := database.DB.First(&dessert, id).Error; err != nil {
		log.Println("错误:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到该甜品"})
		return
	}

	if err := database.DB.Delete(&dessert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "甜品删除成功"})
}
