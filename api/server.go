package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/controllers"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// This is for testing, when done, do well to comment
	seed.Load(server.DB)

	apiPort := os.Getenv("PORT")
	fmt.Printf("Listening to port %s", apiPort)
	if apiPort == "" {
		apiPort = "8089"
	}

	server.Run(":" + apiPort)

}
