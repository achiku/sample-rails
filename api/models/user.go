package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/achiku/sample-rails/api/dao"
	_ "github.com/lib/pq" // postgres driver
)

type User struct {
	ID       int64
	Name     string
	Birthday time.Time
}

func NewUserRepository(q Queryer) *UserRepository {
	return &UserRepository{
		db: q,
	}
}

type UserRepository struct {
	db Queryer
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
	return us, true, nil
}

func (ur *UserRepository) FindByIDRawSQL(ctx context.Context, id int64) (*User, bool, error) {
	var u User
	err := ur.db.QueryRow(`
SELECT
  id
  , name
  , birthday
FROM users
WHERE id = $1`, id).Scan(
		&u.ID,
		&u.Name,
		&u.Birthday,
	)
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
	return us, true, nil
}
