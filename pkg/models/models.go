package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Client struct {
	IdClient    int       `json:"id"`
	ClientName  string    `json:"client_name"`
	ClientMail  string    `json:"client_mail"`
	ClientPass  string    `json:"client_pass"`
	ClientPhone string    `json:"client_phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Progress struct {
	Id             int     `json:"id"`
	Level          int     `json:"level"`
	Points         float32 `json:"points"`
	Tests          int     `json:"tests"`
	CompletedTests int     `json:"completed_tests"`
	AvgPoints      float32 `json:"avg_points"`
	ClientId       int     `json:"client_id"`
}
