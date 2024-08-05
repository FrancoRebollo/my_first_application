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
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error checking database connection: %v", err)
	}

	fmt.Println("Connection successful")

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) UserLogin(userLogin UserLogin) error {
	var us_encrypted_password string
	query := "SELECT us_hash FROM users WHERE us_email = $1"

	row := s.db.QueryRow(query, userLogin.UserEmail)
	err := row.Scan(&us_encrypted_password)

	if err == sql.ErrNoRows {
		// No rows were found, return nil user and no error
		return fmt.Errorf("user not found")
	}

	if err != nil {
		return err
	}

	hashedPassword, err := decrypt(string(us_encrypted_password), os.Getenv("SEED_ENCRIPTATION"))
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userLogin.UserPassword)); err != nil {
		return fmt.Errorf("the given password doesn't match")
	}

	return nil
}

func (s *PostgresStore) UserSignUp(userSignUp UserSignUp) error {
	// Verify the username is not already taken
	existingUser, err := s.UserGetByUserName(userSignUp.UserName)
	if err != nil {
		return fmt.Errorf("error checking username: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("you have to choose another username")
	}
	existingUser, err = s.UserGetByPersonalID(userSignUp.UserPersonalID)

	if err != nil {
		return fmt.Errorf("error checking personal identification: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("personal identification already taken by another person")
	}

	existingUser, err = s.UserGetByEmail(userSignUp.UserEmail)

	if err != nil {
		return fmt.Errorf("error checking user by email: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("email already taken by another person")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userSignUp.UserPassword), bcrypt.DefaultCost)
	encryptedPassword, err := encrypt(string(hashedPassword), os.Getenv("SEED_ENCRIPTATION"))

	if err != nil {
		return err
	}

	sqlStatement := `
		INSERT INTO USERS (US_FIRST_NAME,US_LAST_NAME,US_EMAIL,US_USERNAME,US_HASH,US_PERSONAL_ID,US_BIRTHDAY_DATE,US_CREATE_DATETIME,US_CREATE_USER)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	birthdayDate, err := time.Parse("02-01-2006", userSignUp.UserBirthdayDate)
	if err != nil {
		return fmt.Errorf("error parsing birthday date")
	}

	_, err = s.db.Exec(sqlStatement, userSignUp.UserFirstName, userSignUp.UserLastName,
		userSignUp.UserEmail, userSignUp.UserName, encryptedPassword, userSignUp.UserPersonalID, birthdayDate, time.Now(), "API_USER")

	if err != nil {
		return err
	}
	return nil

}

func (s *PostgresStore) UserGetByUserName(userName string) (*User, error) {
	user := new(User)

	query := "SELECT us_id_user, us_username,us_email FROM users WHERE us_username = $1"

	row := s.db.QueryRow(query, userName)
	err := row.Scan(&user.UserID, &user.UserName, &user.UserEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were found, return nil user and no error
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) UserGetByPersonalID(personalID string) (*User, error) {
	user := new(User)

	query := "SELECT us_id_user, us_username,us_email FROM users WHERE us_personal_id = $1"

	row := s.db.QueryRow(query, personalID)

	err := row.Scan(&user.UserID, &user.UserName, &user.UserEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were found, return nil user and no error
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) UserGetByEmail(email string) (*User, error) {
	user := new(User)

	query := "SELECT us_id_user, us_username,us_email FROM users WHERE us_email = $1"

	row := s.db.QueryRow(query, email)

	err := row.Scan(&user.UserID, &user.UserName, &user.UserEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were found, return nil user and no error
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) UserDelete(int) error {
	return nil
}
