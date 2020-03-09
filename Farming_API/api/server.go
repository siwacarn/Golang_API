package api

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/siwacarn/Golang_API/Farming_API/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("Getting Environmental from file ...")
	}

	server.Initialize(os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	// optional load seed file (mockup data)
	// seed.Load(server.DB)

	// starting services
	server.Run(":8080")
}
