package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"volchok96.com/snippetbox/pkg/models/mysql"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	snippets    *mysql.SnippetModel
	redisClient *redis.Client
	tmplCache   map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "Network address of the web server")
	redisAddr := flag.String("redis", "localhost:6379", "Redis server address") // Added for Redis address
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Set the database password here
	dsn := "my_user:my_password@tcp(mysql)/snippetbox"

	// Create a connection pool for the database
	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Connect to Redis
	rdb, err := connectToRedis(*redisAddr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer rdb.Close()

	// Initialize a new template cache
	tmplCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Add application dependencies
	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		snippets:    &mysql.SnippetModel{DB: db},
		redisClient: rdb, // Storing the Redis client in the application struct
		tmplCache:   tmplCache,
	}

	srv := &http.Server{
		Addr:              *addr,
		ErrorLog:          errorLog,
		Handler:           app.routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
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