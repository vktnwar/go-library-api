package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /books
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// POST /books
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// GET /books/:id
func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}
