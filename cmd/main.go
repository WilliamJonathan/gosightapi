package main

import (
	"fmt"
	"gosightapi/internal/app/bootstrap/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}
	server.StartServer()
}
