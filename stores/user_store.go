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

func (s UserStore) GetRatings(userId string) ([]types.MovieRatingOfUser, error) {
	var ratings []types.MovieRatingOfUser
	query := `SELECT r.id, r.movie_id, r.rating, r.comment, m.name 
FROM movie_ratings r 
INNER JOIN movies m ON m.id=r.movie_id 
WHERE r.user_id = ?`
	rows, err := s.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r types.MovieRatingOfUser
		if err := rows.Scan(&r.Id, &r.MovieId, &r.Rating, &r.Comment, &r.MovieName); err != nil {
			return nil, err
		}
		ratings = append(ratings, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func (s UserStore) GetProfile(id string) (types.Profile, error) {
	var p types.Profile
	if err := s.db.QueryRow("SELECT first_name, last_name, about_me FROM profiles WHERE user_id = ?", id).Scan(&p.FirstName, &p.LastName, &p.AboutMe); err != nil {
		return types.Profile{}, err
	}

	return p, nil
}

func (s UserStore) UpdateProfile(id, firstName, lastName, aboutMe string) error {
	if _, err := s.db.Exec("UPDATE profiles SET first_name = ?, last_name = ?, about_me = ? WHERE user_id = ?", firstName, lastName, aboutMe, id); err != nil {
		return err
	}

	return nil
}
