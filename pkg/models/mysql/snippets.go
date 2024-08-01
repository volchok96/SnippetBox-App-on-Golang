package mysql

import (
	"database/sql"
	"errors"

	"volchok96.com/snippetbox/pkg/models"
)

// Define a type that wraps the sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Method for creating a new note in the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	query := `INSERT INTO latestSnippets (title, content, created, expires)
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
	// SQL query to get one entry
	query := `SELECT id, title, content, created, expires
	FROM latestSnippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(query, id)

	// Initialize the pointer to the new Snippet structure
	snippetSample := &models.Snippet{}

	// Use row.Scan() to copy the values from each field from sql.Row in
	// the corresponding field in the Snippet structure
	// The number of arguments must be exactly the same as the number
	// columns in the database table.
	err := row.Scan(
		&snippetSample.ID,
		&snippetSample.Title,
		&snippetSample.Content,
		&snippetSample.Created,
		&snippetSample.Expires,
	)
	if err != nil {
		// If an error is detected, return our error from the models model.ErrNoRecord.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// If everything is fine, the Snippet object is returned.
	return snippetSample, nil
}

// Method returns the 10 most frequently used notes.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	query := `SELECT id, title, content, created, expires FROM latestSnippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// This postponement statement
	// must be executed *after* checking for an error in the Query() method.
	// If Query() returns an error, it will lead to panic
	// as it will try to close the result set with the value: nil.
	defer rows.Close()

	// Initializing an empty slice to store models objects.Snippets.
	var latestSnippets []*models.Snippet

	// Use rows.Next() to iterate over the result. Provide this method
	// first and then each subsequent record from the database for processing
	// using the rows.Scan() method.
	for rows.Next() {
		// Creating a pointer to the new Snippet structure
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Adding the structure to the slice.
		latestSnippets = append(latestSnippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return latestSnippets, nil
}
