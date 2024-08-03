package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Storage interface {
	UserSignUp(UserSignUp) error
	UserLogin(UserLogin) error
	UserGetByID(int) (*User, error)
	UserDelete(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStrUser := os.Getenv("POSTGRES_USER")
	connStrDBName := os.Getenv("POSTGRES_DBNAME")
	connStrPassword := os.Getenv("POSTGRES_PASSWORD")
	connStrHost := os.Getenv("POSTGRES_HOST")
	connStrPort := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connStrHost, connStrPort, connStrUser, connStrPassword, connStrDBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to Postgre database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error checking database connection: %v", err)
	}

	fmt.Println("Connection successful")

	return &PostgresStore{
		db: db,
	}, nil
}
