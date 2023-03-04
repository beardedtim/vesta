package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetConnection() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	return nil, errors.New("no database connection to get - please ensure you called connect")
}

func Connect() *sql.DB {
	if db == nil {
		postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			Config.Host, Config.Port, Config.Username, Config.Password, Config.Database)

		conn, err := sql.Open("postgres", postgresqlDbInfo)

		if err != nil {
			panic(err)
		}

		db = conn
	} else {
		log.Println("Already connected to the Database. Skipping.")
	}

	// Ensure connected
	err := db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

func Disconnect() error {
	if db != nil {
		return db.Close()
	}

	log.Println("No connection to disconnect from. Skipping")

	return nil
}

func IsHealthy() (bool, error) {
	if db == nil {
		return false, errors.New("no Database connection to check if healthy - please call `Connect` before asking if it is healthy")
	}

	err := db.Ping()

	if err != nil {
		return false, err
	}

	return true, nil
}
