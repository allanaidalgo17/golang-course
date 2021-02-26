package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("Name1")
	stmt.Exec("Name2")

	res, _ := stmt.Exec("Name3")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	rows, _ := res.RowsAffected()
	fmt.Println(rows)
}
