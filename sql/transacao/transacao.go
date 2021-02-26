package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/gocourse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?,?)")

	stmt.Exec(2000, "Name2000")
	stmt.Exec(2001, "Name2001")
	_, err = stmt.Exec(1, "Name1Again") // duplicate key

	// Transaction - aborts if any error occurs
	// in Go, we have to check if there is an error (the program won't stop by itself)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// if here, everything went well (no errors)
	tx.Commit()
}
