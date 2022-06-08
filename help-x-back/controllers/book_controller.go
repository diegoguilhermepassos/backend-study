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
			"error": "cannot find book: " + err.Error(),
		})

		return

	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}
