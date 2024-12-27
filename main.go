package gowordcloudbackend

import (
	"log"
	"net/http"
	"wordcloud/controller"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Register routes
	controller.RegisterWordCloudRoutes(r)

	// Start the server
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
