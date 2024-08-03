package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	UserSignUp(UserSignUp) error
	UserLogin(UserLogin) error
	UserGetByUserName(string) (*User, error)
	UserGetByPersonalID(string) (*User, error)
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
	//fmt.Println(connStr)

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

func (s *PostgresStore) UserSignUp(userSignUp UserSignUp) error {
	_, err := s.UserGetByUserName(userSignUp.userName)

	if err == nil {
		return fmt.Errorf("you have to choose another username")
	}
	_, err = s.UserGetByUserName(userSignUp.userPersonalID)

	if err == nil {
		return fmt.Errorf("personal identification already taken by another person")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userSignUp.userPassword), bcrypt.DefaultCost)
	encryptedPassword, err := encrypt(string(hashedPassword), os.Getenv("SEED_ENCRIPTATION"))

	if err != nil {
		return err
	}

	sqlStatement := `
		INSERT INTO USERS (US_EMAIL,US_USERNAME,US_HASH,US_PERSONAL_ID,US_BIRTHDAY_DATE,US_CREATE_DATETIME,US_CREATE_USER)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err = s.db.Exec(sqlStatement, userSignUp.userEmail, userSignUp.userName, encryptedPassword, userSignUp.userPersonalID, userSignUp.userBirthdayDate, time.Now(), "API_USER")

	if err != nil {
		return err
	}

	return nil

}

func (s *PostgresStore) UserLogin(UserLogin) error {
	return nil
}

func (s *PostgresStore) UserGetByUserName(userName string) (*User, error) {
	user := new(User)

	query := "SELECT us_id_user, us_username,us_email FROM users WHERE us_username = $1"

	row := s.db.QueryRow(query, userName)

	err := row.Scan(&user.userID, &user.userName, &user.userEmail)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) UserGetByPersonalID(personalID string) (*User, error) {
	user := new(User)

	query := "SELECT us_id_user, us_username,us_email FROM users WHERE us_personal_id = $1"

	row := s.db.QueryRow(query, personalID)

	err := row.Scan(&user.userID, &user.userName, &user.userEmail)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) UserDelete(int) error {
	return nil
}
