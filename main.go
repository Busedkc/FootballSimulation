package main

import (
	"database/sql"

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
