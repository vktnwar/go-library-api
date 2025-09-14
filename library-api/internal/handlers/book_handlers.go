package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Lista todos os livros
// @Description Retorna a lista completa de livros cadastrados
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Cria um novo livro
// @Description Cria um livro com título, ISBN e autores opcionais
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Dados do livro"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// GetBook godoc
// @Summary Busca um livro pelo ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Atualiza um livro existente
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Dados atualizados"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [put]
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

// DeleteBook godoc
// @Summary Remove um livro
// @Tags books
// @Param id path int true "Book ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
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
