package models

import (
	"database/sql"
	"errors"
	"time"
)

type Quint struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type QuintModel struct {
	DB *sql.DB
}

func (m *QuintModel) Insert(title string, content string, expires int) (int, error) {

	stmt := `INSERT INTO quints (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (m *QuintModel) Get(id int) (*Quint, error) {
	stmt := `SELECT id, title, content, created, expires FROM quints
			WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Quint{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *QuintModel) Latest() ([]*Quint, error) {
	return nil, nil
}
