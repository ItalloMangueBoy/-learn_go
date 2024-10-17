package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Used from sql.Open() to recognize mysql
)

var url string = "root:mkspop@/devbook?charset=utf8&parseTime=True&loc=Local"

// Connect(): Connect your app to the database
func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
