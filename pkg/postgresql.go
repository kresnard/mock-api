package pkg

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func DBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	log.Println("db connected")
	return db
}
