package models

import (
	"testing"

	txdb "github.com/DATA-DOG/go-txdb"
)

func init() {
	txdb.Register("txdb", "pq", "postgres://rails_todo@localhost:5432/rails_todo?sslmode=disable")
}

// TestMain service package setup/teardonw
func TestMain(m *testing.M) {
}
