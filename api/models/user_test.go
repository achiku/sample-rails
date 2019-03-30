package models

import (
	"context"
	"testing"
	"time"

	"github.com/achiku/sample-rails/api/dao"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestUserRepository_FindByID(t *testing.T) {
	db, cleanup := TestSetupDB(t)
	defer cleanup()

	ctx := boil.SkipHooks(context.Background())
	u1 := TestCreateUserData(ctx, t, db, &dao.User{
		Name:         "achiku",
		Birthday:     time.Date(1985, 8, 18, 0, 0, 0, 0, time.UTC),
		RegisteredAt: time.Now(),
	})
	ur := NewUserRepository(db)
	u, found, err := ur.FindByID(ctx, u1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("want one user")
	}
	t.Logf("%v", u)
}

func TestUserRepository_FindByID_rawsql(t *testing.T) {
	db, cleanup := TestSetupDB(t)
	defer cleanup()

	ctx := boil.SkipHooks(context.Background())
	u1 := TestCreateUserData(ctx, t, db, &dao.User{
		Name:         "achiku",
		Birthday:     time.Date(1985, 8, 18, 0, 0, 0, 0, time.UTC),
		RegisteredAt: time.Now(),
	})
	ur := NewUserRepository(db)
	u, found, err := ur.FindByIDRawSQL(ctx, u1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("want one user")
	}
	t.Logf("%v", u)
}
