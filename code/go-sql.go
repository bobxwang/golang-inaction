package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// import with _: 只导入副作用，也就是说，只执行它的init函数并初始化其中的全局变量 

func main() {
	db, err := sql.Open("mysql", "root:fuck51test@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT mobile FROM user WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var smobile string // we "scan" the result in her

	// Query the id of 13
	err = stmtOut.QueryRow(13).Scan(&smobile) // WHERE id = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %s", smobile)
}
