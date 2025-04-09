package handlers

import (
	"dessert-shop/database"
	"dessert-shop/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetDesserts(c *gin.Context) {
	var desserts []models.Dessert
	var total int64

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc") //desc
	category := c.Query("category")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	query := database.DB.Model(&models.Dessert{}).Preload("Category")
	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	if err := query.Count(&total).Error; err != nil {
		log.Println("统计总数出错:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法统计甜品总数"})
		return
	}

	if err := query.Order(sort + " " + order).Limit(limit).Offset(offset).Find(&desserts).Error; err != nil {
		log.Println("查询甜品出错:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取甜品数据"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"total":    total,
		"desserts": desserts,
	})
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
