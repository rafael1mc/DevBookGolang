package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
		db.Close()
		return nil, err
	}

	return db, nil

}
