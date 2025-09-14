package main

import (
	"library-api/internal/database"
	"library-api/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"

	_ "library-api/docs" // docs gerados pelo swag
)

// @title Library API
// @version 1.0
// @description API para gerenciar livros, autores e empréstimos
// @host localhost:8080
// @BasePath /
func main() {
	// Conecta no banco
	database.Connect()

	// Cria router do Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ajuste conforme front
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rotas para Livros
	books := r.Group("/books")
	{
		books.GET("", handlers.GetBooks)          // GET /books
		books.POST("", handlers.CreateBook)       // POST /books
		books.GET("/:id", handlers.GetBook)       // GET /books/:id
		books.PUT("/:id", handlers.UpdateBook)    // PUT /books/:id
		books.DELETE("/:id", handlers.DeleteBook) // DELETE /books/:id
	}

	// Rotas para Autores
	authors := r.Group("/authors")
	{
		authors.GET("", handlers.GetAuthors)          // GET /authors
		authors.POST("", handlers.CreateAuthor)       // POST /authors
		authors.GET("/:id", handlers.GetAuthor)       // GET /authors/:id
		authors.PUT("/:id", handlers.UpdateAuthor)    // PUT /authors/:id
		authors.DELETE("/:id", handlers.DeleteAuthor) // DELETE /authors/:id
	}

	// Rotas para Empréstimos
	loans := r.Group("/loans")
	{
		loans.GET("", handlers.GetLoans)              // GET /loans
		loans.POST("", handlers.CreateLoan)           // POST /loans
		loans.GET("/:id", handlers.GetLoan)           // GET /loans/:id
		loans.PUT("/:id/return", handlers.ReturnLoan) // PUT /loans/:id/return
		loans.DELETE("/:id", handlers.DeleteLoan)     // DELETE /loans/:id
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Sobe servidor na porta 8080
	r.Run(":8080")
}
