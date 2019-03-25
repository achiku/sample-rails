package models

import (
	"testing"
)

func TestUserRepository_FindByID(t *testing.T) {
	d, cleanup := TestSetupDB(t)
	defer cleanup()

	ur := NewUserRepository(d)
	// ctx := context.Background()
	u, found, err := ur.FindByID(nil, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("want one user")
	}
	t.Logf("%v", u)
}
