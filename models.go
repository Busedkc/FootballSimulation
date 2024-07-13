package main

// Team represents a football team in the league
// it includes the team's ID, name, and strength, which may influence match outcomes
type Team struct {
	ID       int    // unique identifier for the team
	Name     string // name of the team
	Strength int    // numerical representation of the team's strength
}

// match represents a single football match between two teams
// it stores details about the match week, teams involved, goals scored, and the result

type Match struct {
	ID        int    // unique identifier for the match
	Week      int    // week number of the league in which the match occurs
	HomeTeam  int    // ID of the home team
	AwayTeam  int    // ID of the away team
	HomeGoals int    // number of goals scored by the home team
	AwayGoals int    // number of goals scored by the away team
	Result    string // result of the match (e.g., "home_win", "away_win", "draw")
}

// Standing represents the current standing of a team in the league
// It includes wins, losses, draws, and other statistics important for league rankings
type Standing struct {
	TeamID         int // ID of the team
	Played         int //total number of matches played
	Wins           int // total number of matches won
	Draws          int // total number of matches drawn
	Losses         int // total number of matches lost
	GoalsFor       int // total number of goals scored by the team
	GoalsAgainst   int // total number of goals scored against the team
	GoalDifference int // goal difference (GoalsFor - GoalsAgainst)
	Points         int // total points accumulated (3 points for a win, 1 for a draw)
}
