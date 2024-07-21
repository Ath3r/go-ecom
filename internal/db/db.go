package db

import "github.com/go-pg/pg/v10"

var db *pg.DB

func Connect() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
}

func Close() {
	db.Close()
}