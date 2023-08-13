package models

import (
	"database/sql"
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
	return nil, nil
}

func (m *QuintModel) Latest() ([]*Quint, error) {
	return nil, nil
}
