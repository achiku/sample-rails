package models

import (
	"context"
	"testing"

	"github.com/achiku/sample-rails/api/dao"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestCreateUserData(ctx context.Context, t *testing.T, tx Queryer, u *dao.User) *dao.User {
	if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	return u
}
