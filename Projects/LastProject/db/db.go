package db

// we import go-sqlite but we won't use it directly
// we use _ to tell go that we will use this package but won't use it directly
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// initialising a database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	// configure connection pooling = how many open connections we can have to this database
	DB.SetMaxOpenConns(10)
	// how many connection we want to keep open if no one is using any at the moment
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  email TEXT NOT NULL UNIQUE, 
	  password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER,
	  FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	// to execute query statement on DB
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}
