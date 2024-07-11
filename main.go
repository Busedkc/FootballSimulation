package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Team struct {
	ID       int
	Name     string
	Strength int
}

type Match struct {
	ID        int
	Week      int
	HomeTeam  int
	AwayTeam  int
	HomeGoals int
	AwayGoals int
	Result    string
}

type Standing struct {
	TeamID         int
	Played         int
	Wins           int
	Draws          int
	Losses         int
	GoalsFor       int
	GoalsAgainst   int
	GoalDifference int
	Points         int
}

func dbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:çöpadam123@tcp(localhost:3306)/simulation")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func playMatches(db *sql.DB) error {
	//Takımları çek
	teams := []Team{}
	rows, err := db.Query("SELECT id, name, strength FROM Teams")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var t Team
		if err := rows.Scan(&t.ID, &t.Name, &t.Strength); err != nil {
			return err
		}
		teams = append(teams, t)
	}
	//Maçları oynat
	rand.Seed(time.Now().UnixNano())
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

		_, err := db.Exec(`INSERT INTO Matches (week, home_team_id, away_team_id, home_goals, away_goals, match_result) 
		                                        VALUES (?, ?, ?, ?, ?, ?)`,
			1, home.ID, away.ID, homeGoals, awayGoals, result)

		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	//Veritabanı bağlantısını aç
	db, err := dbConn()
	if err != nil {
		fmt.Println("Veritabanı bağlantı hatası:", err)
		return
	}
	defer db.Close()

	//Maçları oynat
	err = playMatches(db)
	if err != nil {
		fmt.Println("Maç oynatma hatası:", err)
	}
}
