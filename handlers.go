package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// getStandings handles HTTP requests to retrieve the current standings from the league table
// It queries the database for all teams' standings and returns them in JSON format
func getStandings(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn() //establishing a database connection
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //ensure the database connection is closed on function exit

	//query the database for all current standing
	rows, err := db.Query("SELECT * FROM Standings")
	if err != nil {
		http.Error(w, "Failed to query standings: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() //ensure rows are closed after processing

	var standings []Standing
	for rows.Next() {
		var s Standing
		//scan the row into the satnding struct
		if err := rows.Scan(&s.TeamID, &s.Played, &s.Wins, &s.Draws, &s.Losses, &s.GoalsFor, &s.GoalsAgainst, &s.GoalDifference, &s.Points); err != nil {
			http.Error(w, "Failed to parse standings: "+err.Error(), http.StatusInternalServerError)
			return
		}
		standings = append(standings, s) //append the parsed standing to the slice
	}

	w.Header().Set("Content-Type", "application/json") //set content type as JSON
	json.NewEncoder(w).Encode(standings)               //encode standings into JSON and send it as a response
}

// playAllMatches handles HTTP requests to simulate all matches for the current week and update the standings.
func playAllMatches(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn() // Establish a database connection.
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() //ensure the database connection is closed on function exit

	//retrieve the current week from the database
	week, err := getCurrentWeek(db)
	if err != nil {
		http.Error(w, "Failed to get current week: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//simulate matches for the current week
	if err := playMatches(db, week); err != nil {
		http.Error(w, "Failed to play matches: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//update the satndings in the database after matches are played
	if err := updateStandings(db); err != nil {
		http.Error(w, "Failed to update standings: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Matches played and standings updated successfully") //send a success message to the client
}
