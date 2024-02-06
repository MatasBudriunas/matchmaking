package main

import (
	"fmt"
	"log"
	"matchmaking/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	fmt.Println("Server is starting on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
