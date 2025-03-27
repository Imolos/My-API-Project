package main

import (
	"log"
	"net/http"

	"github.com/Imolos/My-API-Project/routes"
)

func main() {
	routes.RegisterRoutes()

	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}