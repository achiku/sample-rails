package models

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"testing"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/achiku/sample-rails/api/infra"
	_ "github.com/lib/pq" // postgres driver
)

func init() {
	txdb.Register("txdb", "pq", "postgres://rails_todo_test@localhost:5432/rails_todo_test?sslmode=disable")
}

// TestCreateDB set up test db
func TestCreateDB(path string) error {
	orgPwd, _ := os.Getwd()
	defer func() {
		os.Chdir(orgPwd)
	}()

	os.Chdir(path)
	cmd := exec.Command("rails", "db:create", "RAILS_ENV=test")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("failed to execute rake:\n %s", stderr.String())
		return err
	}
	return nil
}

// TestDropDB set up test db
func TestDropDB(path string) error {
	orgPwd, _ := os.Getwd()
	defer func() {
		os.Chdir(orgPwd)
	}()

	os.Chdir(path)
	cmd := exec.Command("rails", "db:drop", "RAILS_ENV=test")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("failed to execute rake:\n %s", stderr.String())
		return err
	}
	return nil
}

// TestCreateTables create test tables
func TestCreateTables(path string) error {
	orgPwd, _ := os.Getwd()
	defer func() {
		os.Chdir(orgPwd)
	}()

	os.Chdir(path)
	cmd := exec.Command("rails", "db:migrate", "RAILS_ENV=test")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("failed to execute alembic:\n %s", stderr.String())
		return err
	}
	return nil
}

// TestSetupTx create tx and cleanup func for test
func TestSetupTx(t *testing.T) (infra.Txer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		tx.Rollback()
		db.Close()
	}
	return tx, cleanup
}

// TestSetupDB create db and cleanup func for test
func TestSetupDB(t *testing.T) (infra.DBer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		db.Close()
	}
	return db, cleanup
}
