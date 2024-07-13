package main

import "math/rand"

// simulateMatch simulates a football match between two teams and updates the scores and goal differences.
func simulateMatch(match *Match) {
	// Generate random scores for home and away teams based on their strength
	match.HomeScore = rand.Intn(match.HomeTeam.Strength + 1)
	match.AwayScore = rand.Intn(match.AwayTeam.Strength + 1)

	// Update the goal difference based on the match result
	if match.HomeScore > match.AwayScore {
		// Home team wins, update goal differences accordingly
		match.HomeTeam.GoalDifference += match.HomeScore - match.AwayScore
		match.AwayTeam.GoalDifference -= match.HomeScore - match.AwayScore
	} else if match.HomeScore < match.AwayScore {
		// Away team wins, update goal differences accordingly
		match.HomeTeam.GoalDifference -= match.AwayScore - match.HomeScore
		match.AwayTeam.GoalDifference += match.AwayScore - match.HomeScore
	}
}
