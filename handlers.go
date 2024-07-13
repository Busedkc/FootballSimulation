package main

import (
	"encoding/json"
	"net/http"
)

// getTeamsHandler handles the HTTP request for retrieving all teams
func getTeamsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve teams from the database
	teams, err := getTeams()
	if err != nil {
		http.Error(w, "Failed to retrieve teams", http.StatusInternalServerError)
		return
	}
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the teams as JSON
	json.NewEncoder(w).Encode(teams)
}

// getMatchesHandler handles the HTTP request for retrieving all matches
func getMatchesHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve matches from the database
	matches, err := getMatches()
	if err != nil {
		http.Error(w, "Failed to retrieve matches", http.StatusInternalServerError)
		return
	}
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the matches as JSON
	json.NewEncoder(w).Encode(matches)
}

// simulateMatchHandler handles the HTTP request for simulating a match
func simulateMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match Match
	// Decode the request body into the match struct
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Simulate the match
	simulateMatch(&match)
	// Save the match result to the database
	if err := saveMatch(match); err != nil {
		http.Error(w, "Failed to save match", http.StatusInternalServerError)
		return
	}
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the match result as JSON
	json.NewEncoder(w).Encode(match)
}
