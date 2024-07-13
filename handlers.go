package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getStandings(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Standings")
	if err != nil {
		http.Error(w, "Failed to query standings: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var standings []Standing
	for rows.Next() {
		var s Standing
		if err := rows.Scan(&s.TeamID, &s.Played, &s.Wins, &s.Draws, &s.Losses, &s.GoalsFor, &s.GoalsAgainst, &s.GoalDifference, &s.Points); err != nil {
			http.Error(w, "Failed to parse standings: "+err.Error(), http.StatusInternalServerError)
			return
		}
		standings = append(standings, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(standings)
}

func playAllMatches(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := playMatches(db); err != nil {
		http.Error(w, "Failed to play matches: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := updateStandings(db); err != nil {
		http.Error(w, "Failed to update standings: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Matches played and standings updated successfully")
}
