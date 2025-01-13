package server

import (
	"database/sql"
	"fmt"
	"gosightapi/internal/app/bootstrap/database"
	"gosightapi/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	fmt.Println("Starting server...")

	// Conecta ao banco de dados
	db, err := database.DbConnect()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer func() {
		fmt.Println("Closing database connection...")
		db.Close()
	}()

	// Configura o servidor
	e := gin.Default()
	err = configureGroupRoutes(e, db)
	if err != nil {
		fmt.Printf("Failed to configure routes: %v\n", err)
		return
	}

	// Inicializa o servidor na porta 8080
	port := ":8080"
	fmt.Printf("Server is starting on port %s...\n", port)
	err = e.Run(port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}

	fmt.Println("Server is running on port 8080")
}

func configureGroupRoutes(e *gin.Engine, db *sql.DB) error {
	// Define as rotas
	g := e.Group("/api/v1")
	{
		g.GET("/states", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
		g.GET("/items", func(ctx *gin.Context) {
			controllers.GetItems(ctx, db)
		})
	}
	return nil
}
