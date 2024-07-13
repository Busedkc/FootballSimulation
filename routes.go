package main

import "net/http"

// setupRoutes sets up the HTTP routes for the web server
func setupRoutes() {
	// Route for getting all teams
	http.HandleFunc("/teams", getTeamsHandler)

	// Route for getting all matches
	http.HandleFunc("/matches", getMatchesHandler)

	// Route for simulating a match
	http.HandleFunc("/simulate", simulateMatchHandler)
}
