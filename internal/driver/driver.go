package driver

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool
func ConnectSQL(dsn string) (*DB, error) {
	conn, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(maxOpenDbConn)
	conn.SetMaxIdleConns(maxIdleDbConn)
	conn.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = conn

	err = testDB(conn)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// testDB tries to ping the database
func testDB(connection *sql.DB) error {
	err := connection.Ping()
	if err != nil {
		return err

	}
	return nil

}

// NewDatabase creates a new database for the connection pool
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
