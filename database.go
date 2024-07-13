package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// dbConn creates and returns a connection to the database
// If the connection fails, it returns an error
func dbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Busetechista_1@tcp(localhost:3306)/simulation")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// getCurrentWeek retrieves the current week number from the database
// it looks up the 'current_week' key in the Settings table and returns its value
// if the query fails, it returns an error
func getCurrentWeek(db *sql.DB) (int, error) {
	var week int
	err := db.QueryRow("SELECT value FROM Settings WHERE key = 'current_week'").Scan(&week)
	return week, err
}

// updateCurrentWeek updates the current week number in the database
// it sets the value of the 'current_week' key in the Settings table to the specified week number
// if the update fails, it returns an error
func updateCurrentWeek(db *sql.DB, week int) error {
	_, err := db.Exec("UPDATE Settings SET value = ? WHERE key = 'current_week'", week)
	return err
}
