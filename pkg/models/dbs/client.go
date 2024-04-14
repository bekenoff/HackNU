package dbs

import (
	"HackNU/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (m *ClientModel) GetUserById(id int) ([]byte, error) {
	stmt := `SELECT * FROM ainur_hacknu.client WHERE id = ?`
	c := &models.Client{}

	userRow := m.DB.QueryRow(stmt, id)

	err := userRow.Scan(&c.IdClient, &c.ClientMail, &c.ClientPass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedUser, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return convertedUser, nil
}

func (m *ClientModel) GetUserByEmailAndPassword(email, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, nil
	}
	stmt := `SELECT id FROM ainur_hacknu.client WHERE client_mail = ? AND client_pass = ?`
	var userId int

	err = m.DB.QueryRow(stmt, email, hashedPassword).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		} else {
			return 0, err
		}
	}

	return userId, nil
}

func (m *ClientModel) GetLastUserId() (int, error) {
	stmt := `SELECT MAX(id) FROM ainur_hacknu.client`
	var lastId int

	err := m.DB.QueryRow(stmt).Scan(&lastId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		} else {
			log.Printf("Error: %v", err)
			return 0, err
		}
	}

	return lastId, nil
}
