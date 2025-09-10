package handlers

import (
	"library-api/internal/database"
	"library-api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /loans
func GetLoans(c *gin.Context) {
	var loans []models.Loan
	database.DB.Preload("Book.Authors").Find(&loans) // já traz o livro e autores
	c.JSON(http.StatusOK, loans)
}

// POST /loans
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

// GET /loans/:id
func GetLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan

	if err := database.DB.Preload("Book.Authors").First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

// PUT /loans/:id/return
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

// DELETE /loans/:id
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
