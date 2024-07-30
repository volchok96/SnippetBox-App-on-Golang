package mysql

import (
	"database/sql"

	"volchok96.com/snippetbox/pkg/models"
)

// Define a type that wraps the sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Method for creating a new note in the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Method for returning note data by its ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Method returns the 10 most frequently used notes.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
