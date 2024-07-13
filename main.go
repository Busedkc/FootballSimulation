package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Initialize the random number generator
	// This ensures consistency of random numbers used throughout the application
	rand.Seed(time.Now().UnixNano())

	// Set up the HTTP server
	//this defines which endpoints are available and how they are handled
	setupRoutes()

	// Message indicating that the server is starting
	fmt.Println("Server is running on http://localhost:8080")

	//Start the HTTP server and log any errors
	//The 'log.Fatal' function will terminate the application if an error returns from 'http.ListenAndServe'
	//The server starts listening on the specified port (8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
