package main

import (
    "database/sql"
    "fmt"
    "os"
    "strings"

    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/term"
)

// promptForPassword prompts the user for the database password.
func promptForPassword() (string, error) {
    fmt.Print("Enter database password: \n")
    bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
    if err != nil {
        return "", err
    }
    
    // Remove the newline character from the entered password
    password := strings.TrimSpace(string(bytePassword))
    return password, nil
}

// constructDSN constructs the data source name (DSN) string with the entered password.
func constructDSN(password, dsn string) string {
    return fmt.Sprintf("web:%s@%s", password, dsn)
}

// openDB wraps sql.Open() and returns a sql.DB connection pool 
// for the specified data source name (DSN).
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    if err = db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}
