package sqlite

import (
	"database/sql"
	"errors"

	"github.com/Li-Khan/golangify/pkg/models"
)

// SnippetModel ...
type SnippetModel struct {
	DB *sql.DB
}

// Insert ...
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `
	INSERT INTO snippets (
		title, 
		content, 
		created, 
		expires)
    VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, content, "чо?", expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get ...
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}
	err := m.DB.QueryRow("SELECT * FROM snippets WHERE id = ?", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

// Latest ...
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
