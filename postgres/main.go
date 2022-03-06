package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err := sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	cmd := "DELETE FROM persons WHERE name = $1"
	_, err = Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
