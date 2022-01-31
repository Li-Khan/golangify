package sqlite

import (
	"database/sql"

	"github.com/Li-Khan/golangify/pkg/models"
)

// SnippetModel ...
type SnippetModel struct {
	DB *sql.DB
}

// Insert ...
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get ...
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest ...
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
