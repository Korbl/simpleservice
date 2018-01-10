package main


import (
	"log"
	"net/http"
	"simpleservice/handlers"
)

func main() {
	log.Print("Starting the service....")

	router := handlers.Router()
	log.Print("The service is ready to liste and serve.")

	log.Fatal(http.ListenAndServe(":8000", nil))
}