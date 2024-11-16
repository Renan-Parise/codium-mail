package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Renan-Parise/mail/routes"
	"github.com/Renan-Parise/mail/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Is it missing?")
	}

	utils.InitLogger()

	readyChan := make(chan struct{})
	go utils.StartConsumer(readyChan)
	<-readyChan

	router := routes.SetupRouter()
	go func() {
		if err := router.Run("127.0.0.1:8182"); err != nil {
			log.Fatal("Failed to run server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down application...")
}
