package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int

	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) returning id
	`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := `
		SELECT id, username, password, email
		FROM users
		WHERE email = $1
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return &User{}, nil
	}

	return &user, nil
}
