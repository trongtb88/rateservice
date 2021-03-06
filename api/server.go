package api

import (
	"fmt"
	"github.com/trongtb88/rateservice/api/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/trongtb88/rateservice/api/db"
	"github.com/trongtb88/rateservice/api/seed"
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
		fmt.Println("We are getting the env values")
	}
	server.DB = db.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	log.Print("Listening to port 8687")
	seed.Load(server.DB)

	server.Router = mux.NewRouter()

	server.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":8687", server.Router))
}
