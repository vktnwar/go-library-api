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

// PUT /books/:id
func UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var book models.Book

    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    var input models.Book
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// Atualiza dados básicos
	database.DB.Model(&book).Updates(models.Book{
		Title:     input.Title,
		ISBN:      input.ISBN,
		Available: input.Available,
	})

	// Atualiza autores se AuthorIDs foi enviado
	if len(input.AuthorIDs) > 0 {
		var authors []models.Author
		database.DB.Find(&authors, input.AuthorIDs)
		database.DB.Model(&book).Association("Authors").Replace(&authors)
	}

    c.JSON(http.StatusOK, book)
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var book models.Book

    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    database.DB.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
