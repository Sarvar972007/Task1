package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id       int
	UserName string
	Email    string
}
type Storage struct {
	Db *sql.DB
}

type UserFetcher interface {
	GetUserById(ctx context.Context, id int) (*User, error)
}

func (s *Storage) GetUserById(ctx context.Context, id int) (*User, error) {

	query := `SELECT id, username, email FROM users WHERE id = $1`
	row := s.Db.QueryRowContext(ctx, query, id)

	var user User
	err := row.Scan(&user.Id, &user.UserName, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}
	return &user, nil
}
