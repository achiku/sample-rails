package models

import (
	"log"
	"os"
	"testing"
)

// TestMain service package setup/teardonw
func TestMain(m *testing.M) {
	if err := TestDropDB(".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := TestCreateDB(".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := TestCreateTables(".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)

	if err := TestDropDB(".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
