package main

import (
	"log"

	"github.com/Renan-Parise/codium-mail/routes"
	"github.com/Renan-Parise/codium-mail/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Is it missing?")
	}

	utils.InitLogger()

	go utils.StartConsumer()

	router := routes.SetupRouter()
	router.Run("127.0.0.1:8182")
}
