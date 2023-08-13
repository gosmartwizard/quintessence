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
	return 0, nil
}

func (m *QuintModel) Get(id int) (*Quint, error) {
	return nil, nil
}

func (m *QuintModel) Latest() ([]*Quint, error) {
	return nil, nil
}
