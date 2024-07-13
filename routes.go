package main

import "net/http"

// setupRoutes configures the HTTP routes for the application
// It maps specific endpoints to their corresponding handler functions
func setupRoutes() {
	// Handle HTTP requests to "/standings" with the getStandings function
	// This endpoint provides the current standings in the league
	http.HandleFunc("/standings", getStandings)

	// Handle HTTP requests to "/playall" with the playAllMatches function
	// this endpoint simulates playing all matches for the current week and updates the standings accordingly
	http.HandleFunc("/playall", playAllMatches)

	// Handle HTTP requests to "/advanceWeek"
	//this endpoint advances the league to the next week, updating all relevant matches and standings
	http.HandleFunc("/advanceWeek", func(w http.ResponseWriter, r *http.Request) {
		db, err := dbConn() //establish a database connection.
		if err != nil {
			//return an error message if the database connection fails
			http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close() //ensure the database connection is closed on function exit

		// try to advance the week and handle any errors
		if err := advanceWeek(db); err != nil {
			//return an error message if advancing the week fails
			http.Error(w, "Error advancing week: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// if the week is advanced successfully, send a success message
		w.Write([]byte("Week advanced successfully."))
	})
}
