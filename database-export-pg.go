package db_export

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func connectDatabase() {
	host := "localhost"
	port := 5434
	user := "postgres"
	pass := "pgpass"
	dbname := "test_db"

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("error during database connection", errSql)
	} else {
		Db = db
		fmt.Println("connected sucessfully")
	}

}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
