package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ = sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	/*
		cmd := `CREATE TABLE IF NOT EXISTS persons(
						name STRING,
						age  INT)`
		_, err := Db.Exec(cmd)
	*/

	/*
		cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err := Db.Exec(cmd, "tarou", 20)
	*/

	/*
		cmd := "UPDATE persons SET age = ? WHERE name = ?"
		_, err := Db.Exec(cmd, 25, "tarou")
	*/

	/*
		cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err := Db.Exec(cmd, "hanako", 19)
	*/

	/*
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		cmd := "SELECT * FROM persons where age = ?"
		row := Db.QueryRow(cmd, 25)
		var p Person
		err := row.Scan(&p.Name, &p.Age)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Println("err")
			}
		}

		fmt.Println(p.Name, p.Age)
	*/

	/*
		cmd := "SELECT * FROM persons"
		rows, _ := Db.Query(cmd)
		defer rows.Close()

		var pp []Person
		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Println(err)
			}
			pp = append(pp, p)
		}

		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	cmd := "DELETE FROM  persons WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")

	if err != nil {
		log.Fatalln(err)
	}

}
