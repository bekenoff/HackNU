package main

import (
	"HackNU/pkg/models"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type ClientModel struct {
	DB *sql.DB
}

func (m *ClientModel) Insert(clientmail, clientpass string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clientpass), 12)
	if err != nil {
		return err
	}

	stmt := `
        INSERT INTO ainur_hacknu.client 
        (client_mail, client_pass) 
        VALUES (?, ?);`

	_, err = m.DB.Exec(stmt, clientmail, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (m *ClientModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, client_pass FROM ainur_hacknu.client  WHERE client_mail = ?"
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}
