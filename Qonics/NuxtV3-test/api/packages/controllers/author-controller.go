package controllers

import (
	"github.com/dally469/api/packages/config"
	"github.com/dally469/api/packages/helper"
	"github.com/dally469/api/packages/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ManipulateAuthor(c *gin.Context) {

	helper.RequestAppendHeader(c)
	var authors  models.Author
	if err := c.ShouldBindJSON(&authors); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid data" + err.Error()})
		return
	}
	authors.Status = 0
	authors.ID = uuid.NewV4().String()

	if err := config.DB.Create(&authors).Error; err != nil {
		c.JSON(400, gin.H{"status": 500, "message": "Create item failed, " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Author created successfully", "data": authors})
}
