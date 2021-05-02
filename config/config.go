package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetDb() (*sql.DB, error){
	db, err := sql.Open("sqlite3", "db/book.db")
	return db, err
}