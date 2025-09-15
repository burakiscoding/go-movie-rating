package stores

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/google/uuid"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) UserStore {
	return UserStore{
		db: db,
	}
}

func (s UserStore) CreateUserAndProfile(ctx context.Context, email, password string) (string, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var userExists bool
	if err = tx.QueryRowContext(ctx, "SELECT (email = ?) FROM users WHERE email = ?", email, email).Scan(&userExists); err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if userExists {
		return "", fmt.Errorf("User already exists")
	}

	id := uuid.New().String()

	_, err = tx.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES (?, ?, ?)", id, email, password)
	if err != nil {
		return "", err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO profiles (first_name, last_name, about_me, user_id) VALUES (?, ?, ?, ?)", "", "", "", id)
	if err != nil {
		return "", err
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}

	return id, nil
}

func (s UserStore) GetUserByEmail(email string) (types.User, error) {
	var user types.User
	if err := s.db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return types.User{}, err
	}

	return user, nil
}
