package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Busetechista_1@tcp(localhost:3306)/simulation")
	if err != nil {
		return nil, err
	}
	return db, nil
}
