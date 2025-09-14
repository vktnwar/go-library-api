package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetLoans godoc
// @Summary Lista todos os empréstimos
// @Tags loans
// @Produce json
// @Success 200 {array} models.Loan
// @Router /loans [get]
func GetLoans(c *gin.Context) {
	var loans []models.Loan
	database.DB.Preload("Book.Authors").Find(&loans) // já traz o livro e autores
	c.JSON(http.StatusOK, loans)
}

// CreateLoan godoc
// @Summary Cria um novo empréstimo
// @Tags loans
// @Accept json
// @Produce json
// @Param loan body models.Loan true "Dados do empréstimo"
// @Success 201 {object} models.Loan
// @Failure 400 {object} map[string]string
// @Router /loans [post]
func CreateLoan(c *gin.Context) {
	var loan models.Loan

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o livro existe e está disponível
	var book models.Book
	if err := database.DB.First(&book, loan.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if !book.Available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book is not available"})
		return
	}

	// Marca livro como indisponível
	book.Available = false
	database.DB.Save(&book)

	// Define data do empréstimo
	loan.LoanDate = time.Now()

	database.DB.Create(&loan)
	c.JSON(http.StatusCreated, loan)
}

// GetLoan godoc
// @Summary Busca um empréstimo pelo ID
// @Tags loans
// @Produce json
// @Param id path int true "Loan ID"
// @Success 200 {object} models.Loan
// @Failure 404 {object} map[string]string
// @Router /loans/{id} [get]
func GetLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan

	if err := database.DB.Preload("Book.Authors").First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

// ReturnLoan godoc
// @Summary Marca um empréstimo como devolvido
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Success 200 {object} models.Loan
// @Failure 404 {object} map[string]string
// @Router /loans/{id}/return [put]
func ReturnLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan

	if err := database.DB.Preload("Book").First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	// Se já foi devolvido
	if loan.ReturnDate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book already returned"})
		return
	}

	// Atualiza data de devolução
	now := time.Now()
	loan.ReturnDate = &now
	database.DB.Save(&loan)

	// Marca livro como disponível novamente
	loan.Book.Available = true
	database.DB.Save(&loan.Book)

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully", "loan": loan})
}

// DeleteLoan godoc
// @Summary Remove um empréstimo
// @Tags loans
// @Param id path int true "Loan ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /loans/{id} [delete]
func DeleteLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan

	if err := database.DB.Preload("Book").First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	// Se empréstimo ainda estava ativo, libera o livro
	if loan.ReturnDate == nil {
		loan.Book.Available = true
		database.DB.Save(&loan.Book)
	}

	database.DB.Delete(&loan)
	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
}
