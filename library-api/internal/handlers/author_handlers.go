package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /authors
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	database.DB.Preload("Books").Find(&authors) // já traz livros
	c.JSON(http.StatusOK, authors)
}

// POST /authors
func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

// GET /authors/:id
func GetAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author

	if err := database.DB.Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// PUT /authors/:id
func UpdateAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author

	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	var input models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&author).Updates(models.Author{
		Name: input.Name,
		Bio:  input.Bio,
	})

	c.JSON(http.StatusOK, author)
}

// DELETE /authors/:id
func DeleteAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author

	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	database.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
