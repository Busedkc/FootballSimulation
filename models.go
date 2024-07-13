package main

// Team represents a football team with relevant attributes.
type Team struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Strength       int    `json:"strength"`
	GoalDifference int    `json:"goal_difference"`
}

// Match represents a football match between two teams with the respective scores.
type Match struct {
	HomeTeam  Team `json:"home_team"`
	AwayTeam  Team `json:"away_team"`
	HomeScore int  `json:"home_score"`
	AwayScore int  `json:"away_score"`
}

// League represents a collection of teams and matches.
type League struct {
	Teams   []Team  `json:"teams"`
	Matches []Match `json:"matches"`
}
