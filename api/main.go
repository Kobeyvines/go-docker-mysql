package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connect to the MySQL database
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", "root", "rootpassword", "db", "testdb")
	db, err = sql.Open("mysql", dsn)
	if err 

}
