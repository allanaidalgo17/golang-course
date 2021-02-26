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

	// update
	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("New Name1", 1)
	stmt.Exec("New Name2", 2)

	// delete
	stmt2, _ := db.Prepare("delete from users where id = ?")
	stmt2.Exec(3)
}
