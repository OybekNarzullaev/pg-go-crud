package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var DBClient *sqlx.DB

func InitialDBConnection() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	conStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName)

	db, err := sqlx.Open("postgres", conStr)
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	DBClient = db
}
