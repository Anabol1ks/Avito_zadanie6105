package main

import (
	"log"
	"os"
	"tender/db"
	"tender/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	db.ConDB()

	router := routes.SetupRouter()

	serverAddress := os.Getenv("SERVER_ADDRESS")

	router.Run(serverAddress)
}
