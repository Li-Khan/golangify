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
	stmt := `SELECT * FROM snippets WHERE expires = 7 ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
