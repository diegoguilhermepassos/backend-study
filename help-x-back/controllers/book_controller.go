package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperguri/webapi-with-go/database"
	"github.com/hyperguri/webapi-with-go/models"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has no to be integer",
		})
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot fin book: " + err.Error(),
		})

		return

	}
	c.JSON(200, book)
}
