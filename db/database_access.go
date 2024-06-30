package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// the sql.Open is used to initialize a connection to a DB -> driver name , path to a DB
	DB, err = sql.Open("sqlite3", "gopham_general_db.db")

	if err != nil {
		panic("error in connecting to database")
	}

	// to configure a connection pool
	DB.SetMaxOpenConns(10)
	// how many connection to keep open if they are not being used
	DB.SetConnMaxIdleTime(5)
	createTables()
}

func createTables() {
	createPatientsTable := `create table if not exists patients (
    	id integer primary key autoincrement,
    	first_name text not null,
    	last_name text not null,
    	is_insured integer,
    	diseases text)`

	result, err := DB.Exec(createPatientsTable)
	if err != nil {
		panic("couldn't create patients table")
	}

	fmt.Println("result of patients table creation execution", result)

	createStaffTable := `create table if not exists staff (
    	id integer primary key autoincrement,
    	first_name text not null,
    	last_name text not null,
    	password text unique not null,
    	job text,
    	role text)`

	result, err = DB.Exec(createStaffTable)
	if err != nil {
		panic("couldn't create staff table")
	}

	fmt.Println("result of staff table creation execution", result)
}
