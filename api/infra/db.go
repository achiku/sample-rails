package infra

import (
	"context"
	"database/sql"
)

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Txer database/sql transaction interface
type Txer interface {
	Queryer
	Commit() error
	Rollback() error
}

// DBer database/sql
type DBer interface {
	Queryer
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}
