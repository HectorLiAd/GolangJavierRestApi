package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionStr := "root:hector@tcp(localhost:3306)/northwind"
	fmt.Println(connectionStr)
	databaseConnection, err := sql.Open("mysql", connectionStr)

	if err != nil || databaseConnection == nil {
		panic(err.Error()) //Error Handling manejo de errores
	}
	return databaseConnection
}
