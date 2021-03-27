package api

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/trongtb88/rateservice/api/db"
	"github.com/trongtb88/rateservice/api/seed"
)


type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

var server = Server{}


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
	log.Print("Listening to port 8686")
	log.Print("Go to here")
	seed.Load(server.DB)

	server.Router = mux.NewRouter()

	server.initializeRoutes()

	log.Fatal(http.ListenAndServe(":8686", server.Router))
}
