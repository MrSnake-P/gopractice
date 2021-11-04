package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { _ = db.Close() }()
	row := db.QueryRow("SELECT  NAME  FROM USER LIMIT 1")
	var name string
	if err := row.Scan(&name); err != nil {
		fmt.Println(err)
	}
	log.Println(name)

}