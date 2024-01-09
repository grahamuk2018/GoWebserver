package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "6543"
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func GetConnection() *sqlx.DB {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlinfo)

	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")
	return db
}
