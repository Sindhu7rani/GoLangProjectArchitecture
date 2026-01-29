package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
Create database connection
DB connection is created once when the application starts
and reused for all incoming requests
*/
func ConnectDB() (*sql.DB, error) {
	/*creating database in given DB driver
	here,we r creating golangArche DB in sqlite */
	db, err := sql.Open("sqlite3", "./golangArche.db")
	//if err while creting DB -->err
	if err != nil {
		return nil, err
	}
	//checking if DB is accessable, if not-->err
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//creating tables
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		balance REAL
	);
	`
	//excuting table creation query
	_, err = db.Exec(query)
	//table not created-->err
	if err != nil {
		return nil, err
	}
	//if DB created successfully send DB
	return db, nil
}
