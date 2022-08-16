package controllers

import (
	// "time"

	
	"github.com/dally469/api/packages/config"
	"github.com/dally469/api/packages/helper"
	"github.com/dally469/api/packages/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)




func CreateBook(c *gin.Context) {
	helper.RequestAppendHeader(c)
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid data" + err.Error()})
		return
	}
	book.Status = 1
	book.Id = uuid.NewV4().String()

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(400, gin.H{"status": 500, "message": "Create item failed, " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Row created successfully", "data": book})
}

func GetAllBook(c *gin.Context) {
	helper.RequestAppendHeader(c)
	var books []models.Book
	if err := config.DB.Select("*").
		Find(&books).Error; err != nil {
		c.JSON(404, gin.H{"status": 400, "message": "No periods found " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "books": books})
}

func GetBoodById(c *gin.Context)  {
	helper.RequestAppendHeader(c)
	var books []models.Book
	bookId := c.Param("id")
	 if err := config.DB.First(&books, "id = ?", bookId).
	 Error; err != nil {
			c.JSON(404, gin.H{"status": 400, "message": "No Book found " + err.Error()})
			return
	 }
	 c.JSON(200, gin.H{"status":200, "book": books})
}