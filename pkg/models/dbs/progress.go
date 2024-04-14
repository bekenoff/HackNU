package dbs

import (
	"HackNU/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
)

type ProgressModel struct {
	DB *sql.DB
}

func (m *ProgressModel) Insert(progress *models.Progress) error {
	stmt := `
        INSERT INTO ainur_hacknu.progress 
        ( level, points, tests, films, meetings, client_id) 
        VALUES (?, ?, ?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, progress.Level, progress.Points, progress.Tests, progress.Films, progress.Meetings, progress.ClientId)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) GetProgressById(id int) ([]byte, error) {
	stmt := `SELECT * FROM ainur_hacknu.progress WHERE client_id = ?`

	progressRow := m.DB.QueryRow(stmt, id)

	p := &models.Progress{}

	err := progressRow.Scan(&p.Id, &p.Level, &p.Points, &p.Tests, &p.Films, &p.Meetings, &p.ClientId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedProgress, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return convertedProgress, nil
}

func (m *ProgressModel) UpdateLevel(id string) error {

	var points int
	query := `SELECT points FROM ainur_hacknu.progress WHERE client_id = ?`
	err := m.DB.QueryRow(query, id).Scan(&points)
	if err != nil {
		return err
	}

	var level string
	switch {
	case points < 3:
		level = "A1"
	case points >= 3 && points < 50:
		level = "A2"
	case points >= 50 && points < 100:
		level = "B1"
	case points >= 100 && points < 200:
		level = "B2"
	case points >= 200 && points < 400:
		level = "C1"
	case points >= 400:
		level = "C2"
	}

	stmt := `UPDATE ainur_hacknu.progress SET level = ? WHERE client_id = ?`
	_, err = m.DB.Exec(stmt, level, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) UpdatePoints(id string) error {
	stmt := `UPDATE ainur_hacknu.progress SET points = points + 1 WHERE client_id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = m.UpdateLevel(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) UpdateTests(id string) error {
	stmt := `UPDATE ainur_hacknu.progress SET tests = tests + 1 WHERE client_id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = m.UpdatePoints(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) UpdateFilms(id string) error {
	stmt := `UPDATE ainur_hacknu.progress SET films = films + 1 WHERE client_id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = m.UpdatePoints(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) UpdateMeetings(id string) error {
	stmt := `UPDATE ainur_hacknu.progress SET meetings = meetings + 1 WHERE client_id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = m.UpdatePoints(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ProgressModel) DeleteProgressById(id int) error {
	stmt := `DELETE FROM ainur_hacknu.progress WHERE client_id = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
