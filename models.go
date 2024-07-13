package main

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
