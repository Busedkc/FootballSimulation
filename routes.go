package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("/standings", getStandings)
	http.HandleFunc("/playall", playAllMatches)
}
