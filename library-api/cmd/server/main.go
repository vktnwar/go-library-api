package main

import (
	"library-api/internal/database"
	"library-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta no banco
	database.Connect()

	// Cria router do Gin
	r := gin.Default()

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

	// Sobe servidor na porta 8080
	r.Run(":8080")
}
