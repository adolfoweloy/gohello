package main

import (
	"log"
	"net/http"
	"os"

	"adolfoeloy.com/gohello/handlers"
	"adolfoeloy.com/gohello/services"
)

func main() {
	// Initialize the URL shortener service
	shortener := services.NewURLShortener()

	// Initialize handlers with the service
	handler := handlers.NewHandler(shortener)

	// Set up routes
	http.HandleFunc("/", handler.ServeHomePage)
	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/api/urls", handler.ListURLs)
	http.HandleFunc("/api/stats/", handler.GetStats)
	http.HandleFunc("/redirect/", handler.RedirectURL)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("URL Shortener starting on port %s", port)
	log.Printf("Visit http://localhost:%s to get started", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
