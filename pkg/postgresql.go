package pkg

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func DBConnection() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/mock_api")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("db connected")
	DB = db
}
