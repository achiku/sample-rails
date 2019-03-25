package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/achiku/sample-rails/api/dao"
	"github.com/achiku/sample-rails/api/infra"
	_ "github.com/lib/pq" // postgres driver
)

type User struct {
	ID       int64
	Name     string
	Birthday time.Time
}

func NewUserRepository(q infra.Queryer) *UserRepository {
	return &UserRepository{
		db: q,
	}
}

type UserRepository struct {
	db infra.Queryer
}

func (ur *UserRepository) FindByID(ctx context.Context, id int64) (*User, bool, error) {
	u, err := dao.FindUser(ctx, ur.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}
	us := &User{
		ID:       u.ID,
		Name:     u.Name,
		Birthday: u.Birthday,
	}
	return us, false, nil
}
