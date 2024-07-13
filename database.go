package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initDB initializes the database connection
func initDB() {
	var err error
	// Open a database connection
	db, err = sql.Open("mysql", "root:Busetechista_1@tcp(localhost:3306)/simulation")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verify the connection is established
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database connection established")
}

// getTeams retrieves all teams from the database
func getTeams() ([]Team, error) {
	// Execute the query
	rows, err := db.Query("SELECT id, name, strength, goal_difference FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []Team
	// Iterate over the result set
	for rows.Next() {
		var team Team
		// Scan the row into the team struct
		if err := rows.Scan(&team.ID, &team.Name, &team.Strength, &team.GoalDifference); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

// getMatches retrieves all matches from the database
func getMatches() ([]Match, error) {
	// Execute the query
	rows, err := db.Query("SELECT home_team_id, away_team_id, home_score, away_score FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []Match
	// Iterate over the result set
	for rows.Next() {
		var match Match
		var homeTeamID, awayTeamID int
		// Scan the row into the match struct
		if err := rows.Scan(&homeTeamID, &awayTeamID, &match.HomeScore, &match.AwayScore); err != nil {
			return nil, err
		}
		// Retrieve team details
		match.HomeTeam, _ = getTeamByID(homeTeamID)
		match.AwayTeam, _ = getTeamByID(awayTeamID)
		matches = append(matches, match)
	}
	return matches, nil
}

// getTeamByID retrieves a team by its ID
func getTeamByID(id int) (Team, error) {
	var team Team
	// Execute the query and scan the result into the team struct
	err := db.QueryRow("SELECT id, name, strength, goal_difference FROM teams WHERE id = ?", id).Scan(&team.ID, &team.Name, &team.Strength, &team.GoalDifference)
	if err != nil {
		return team, err
	}
	return team, nil
}

// saveMatch saves a match result into the database
func saveMatch(match Match) error {
	// Execute the insert statement
	_, err := db.Exec("INSERT INTO matches (home_team_id, away_team_id, home_score, away_score) VALUES (?, ?, ?, ?)", match.HomeTeam.ID, match.AwayTeam.ID, match.HomeScore, match.AwayScore)
	return err
}
