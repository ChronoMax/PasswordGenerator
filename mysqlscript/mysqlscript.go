package mysqlscript

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertToDatabase(result string) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("Cannot find MSQL database!")
	}
	defer db.Close()

	// prepare a statement for inserting data
	stmt, err := db.Prepare("INSERT INTO passwords(passwords) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// execute the statement with parameter values
	stmt.Exec(result)
}
