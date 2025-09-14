package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAuthors godoc
// @Summary Lista todos os autores
// @Tags authors
// @Produce json
// @Success 200 {array} models.Author
// @Router /authors [get]
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	database.DB.Preload("Books").Find(&authors) // já traz livros
	c.JSON(http.StatusOK, authors)
}

// CreateAuthor godoc
// @Summary Cria um novo autor
// @Tags authors
// @Accept json
// @Produce json
// @Param author body models.Author true "Dados do autor"
// @Success 201 {object} models.Author
// @Failure 400 {object} map[string]string
// @Router /authors [post]
func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

// GetAuthor godoc
// @Summary Busca um autor pelo ID
// @Tags authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [get]
func GetAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author

	if err := database.DB.Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// UpdateAuthor godoc
// @Summary Atualiza um autor
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Dados atualizados"
// @Success 200 {object} models.Author
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [put]
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

// DeleteAuthor godoc
// @Summary Remove um autor
// @Tags authors
// @Param id path int true "Author ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [delete]
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
