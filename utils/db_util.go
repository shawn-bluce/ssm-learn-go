package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:22c3683b094136c3@tcp(127.0.0.1:3306)/SSM")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
