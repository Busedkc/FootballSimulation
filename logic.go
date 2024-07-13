package main

import (
	"database/sql"
	"math/rand"
)

// playMatches takes a database connection and a week number,
// and simulates matches for that week based on team strengths
func playMatches(db *sql.DB, week int) error {
	//query the database for all teams and their strengths
	rows, err := db.Query("SELECT id, name, strength FROM Teams")
	if err != nil {
		return err
	}
	defer rows.Close()

	teams := []Team{}
	for rows.Next() {
		var t Team
		if err := rows.Scan(&t.ID, &t.Name, &t.Strength); err != nil {
			return err
		}
		teams = append(teams, t)
	}

	//simulate matches for each pair of teams
	for i := 0; i < len(teams); i += 2 {
		home := teams[i]
		away := teams[i+1]

		homeGoals := rand.Intn(home.Strength + 1)
		awayGoals := rand.Intn(away.Strength + 1)
		result := "draw"
		if homeGoals > awayGoals {
			result = "home_win"
		} else if homeGoals < awayGoals {
			result = "away_win"
		}

		//Insert the match result into the Matches table
		_, err := db.Exec(`INSERT INTO Matches (week, home_team_id, away_team_id, home_goals, away_goals, match_result) 
		                                        VALUES (?, ?, ?, ?, ?, ?)`,
			1, home.ID, away.ID, homeGoals, awayGoals, result)

		if err != nil {
			return err
		}
	}
	return nil
}

// updateStandings updates the league standings based on the results of matches
func updateStandings(db *sql.DB) error {
	//Query for all match results
	rows, err := db.Query("SELECT home_team_id, away_team_id, home_goals, away_goals FROM Matches")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var homeID, awayID, homeGoals, awayGoals int
		if err := rows.Scan(&homeID, &awayID, &homeGoals, &awayGoals); err != nil {
			return err
		}

		//determine the result and points for home and away teams
		homeResult, awayResult := "draw", "draw"
		homePoints, awayPoints := 1, 1

		if homeGoals > awayGoals {
			homeResult, awayResult = "win", "loss"
			homePoints, awayPoints = 3, 0
		} else if homeGoals < awayGoals {
			homeResult, awayResult = "loss", "win"
			homePoints, awayPoints = 3, 0
		}
		//update standings for both teams
		if err := updateTeamStanding(db, homeID, homeGoals, awayGoals, homePoints, homeResult); err != nil {
			return err
		}

		if err := updateTeamStanding(db, awayID, awayGoals, homeGoals, awayPoints, awayResult); err != nil {
			return err
		}
	}
	return nil
}

// updateTeamStanding updates a single team's standing in the league table
func updateTeamStanding(db *sql.DB, teamID, goalsFor, goalsAgainst, points int, result string) error {
	_, err := db.Exec(`UPDATE Standings SET
			               played = played + 1,
			               wins = wins + CASE WHEN ? = 'win' THEN 1 ELSE 0 END,
			               draws = draws + CASE WHEN ? = 'draw' THEN 1 ELSE 0 END,
			               losses = losses + CASE WHEN ? = 'loss' THEN 1 ELSE 0 END,
			               goals_for = goals_for + ?,
			               goals_against = goals_against + ?,
			               goal_difference = goal_difference + (? - ?),
			               points = points + ?
			               WHERE team_id = ?`,
		result, result, result, goalsFor, goalsAgainst, goalsFor, goalsAgainst, points, teamID)
	return err
}

// advanceWeek advances the week number in the simulation,
// triggering the playing of matches and updating of standings
func advanceWeek(db *sql.DB) error {
	//retrieve the current week number from the database
	week, err := getCurrentWeek(db)
	if err != nil {
		return err
	}

	// Simulate matches for the current week
	if err := playMatches(db, week); err != nil {
		return err
	}

	// Increment the week number
	week++
	if err := updateCurrentWeek(db, week); err != nil {
		return err
	}

	return nil
}
