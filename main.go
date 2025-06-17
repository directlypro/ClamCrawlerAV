package main

import (
	"clam_crawler_av/internal/http_server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Starting ClamCrawlerAV Service")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Port Configuration
	port := os.Getenv("PORT")
	if port == "" || port == "-" {
		port = "8050"
		log.Printf("Defaulting to port %s", port)
	}

	appRouter := http_server.NewServiceRouter()

	//TODO: Add a timeout for larger file uploads
	server := &http.Server{
		Addr:    ":" + port,
		Handler: appRouter,
	}

	fmt.Printf("Server Listening on port %s", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %+v", err)
	}

}
