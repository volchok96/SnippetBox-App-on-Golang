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
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}
	//  LastInsertId() - method from sql.Result interface
	// returns int64
	// DO NOT USE IT IN POSTGRESQL !
	id, err := result.LastInsertId() 
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

// Method for returning note data by its ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Method returns the 10 most frequently used notes.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
