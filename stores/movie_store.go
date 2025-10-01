package stores

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/burakiscoding/go-movie-rating/types"
)

type MovieStore struct {
	db *sql.DB
}

func NewMovieStore(db *sql.DB) MovieStore {
	return MovieStore{
		db: db,
	}
}

func (s MovieStore) GetAll() ([]types.Movie, error) {
	var movies []types.Movie
	rows, err := s.db.Query("SELECT id, name, release_date, duration_in_minutes, rating, (SELECT name FROM movie_medias WHERE movie_id=movies.id AND type=2 LIMIT 1) AS poster FROM movies")
	if err != nil {
		return nil, fmt.Errorf("MovieStore GetAll Error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m types.Movie
		if err := rows.Scan(&m.Id, &m.Name, &m.ReleaseDate, &m.DurationInMinutes, &m.Rating, &m.Poster); err != nil {
			return nil, fmt.Errorf("MovieStore GetAll Error: %v", err)
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MovieStore GetAll Error: %v", err)
	}

	return movies, nil
}

func (s MovieStore) GetById(id int) (types.MovieDetail, error) {
	var movie types.MovieDetail
	query := `SELECT movies.id, movies.name, movies.release_date, movies.duration_in_minutes, movies.rating, movies.description,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'name', name)) 
FROM
(SELECT DISTINCT genres.id, genres.name
FROM genres INNER JOIN movie_genres ON genres.id=movie_genres.genre_id WHERE movie_genres.movie_id=movies.id) AS t1) AS genres,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'firstName', first_name,'lastName',  last_name)) 
FROM
(SELECT DISTINCT actors.id, actors.first_name, actors.last_name 
FROM actors INNER JOIN movie_actors ON actors.id=movie_actors.actor_id WHERE movie_actors.movie_id=movies.id) AS t2) AS actors,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('name', name, 'type', type))
FROM
(SELECT DISTINCT movie_medias.name, movie_medias.type 
FROM movie_medias WHERE movie_medias.movie_id=movies.id) AS t3) AS medias
FROM movies WHERE id = ?`
	if err := s.db.QueryRow(query, id).Scan(&movie.Id, &movie.Name, &movie.ReleaseDate, &movie.DurationInMinutes, &movie.Rating, &movie.Description, &movie.Genres, &movie.Actors, &movie.Medias); err != nil {
		return types.MovieDetail{}, err
	}

	return movie, nil
}

func (s MovieStore) AddFile(movieId int, fileName string, fileType uint8) error {
	if _, err := s.db.Exec("INSERT INTO movie_medias (name, type, movie_id) VALUES (?, ?, ?)", fileName, fileType, movieId); err != nil {
		return err
	}

	return nil
}

func (s MovieStore) AddRating(ctx context.Context, userId string, movieId int, rating float64, comment string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Insert rating
	_, err = tx.ExecContext(ctx, "INSERT INTO movie_ratings (user_id, movie_id, rating, comment) VALUES (?, ?, ?, ?)", userId, movieId, rating, comment)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Read movie's current rating and number of ratings
	var numberOfRatings int
	var currrentRating float64
	if err := tx.QueryRowContext(ctx, "SELECT number_of_ratings, rating FROM movies WHERE id = ?", movieId).Scan(&numberOfRatings, &currrentRating); err != nil {
		fmt.Println(err)
		return err
	}

	newAverage := CalculateAverage(numberOfRatings, currrentRating, rating)

	// Update movie's rating and increase by one the number of ratings
	if _, err := tx.ExecContext(ctx, "UPDATE movies SET rating = ?, number_of_ratings = number_of_ratings + 1 WHERE id = ?", newAverage, movieId); err != nil {
		fmt.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("success")

	return nil
}

func (s MovieStore) GetRatings(id int) ([]types.MovieRating, error) {
	var movieRatings []types.MovieRating
	query := `SELECT m.id, m.movie_id, m.rating, m.comment, p.first_name, p.last_name
FROM movie_ratings m 
INNER JOIN profiles p ON p.user_id=m.user_id 
WHERE m.movie_id = ?`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m types.MovieRating
		if err := rows.Scan(&m.Id, &m.MovieId, &m.Rating, &m.Comment, &m.FirstName, &m.LastName); err != nil {
			return nil, err
		}
		movieRatings = append(movieRatings, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movieRatings, nil
}

func (s MovieStore) IsMovieExists(id int) bool {
	var isExists bool
	if err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM movies WHERE id = ?) AS is_exist", id).Scan(&isExists); err != nil {
		return false
	}

	return isExists
}
