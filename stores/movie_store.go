package stores

import (
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
	rows, err := s.db.Query("SELECT id, name, release_date, duration_in_minutes, rating FROM movies")
	if err != nil {
		return nil, fmt.Errorf("MovieStore GetAll Error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m types.Movie
		if err := rows.Scan(&m.Id, &m.Name, &m.ReleaseDate, &m.DurationInMinutes, &m.Rating); err != nil {
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
	query := `SELECT movies.id, movies.name, 
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'name', name)) 
FROM
(SELECT DISTINCT genres.id, genres.name
FROM genres INNER JOIN movie_genres ON genres.id=movie_genres.genre_id WHERE movie_genres.movie_id=movies.id) AS t1) AS genres,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'firstName', first_name,'lastName',  last_name)) 
FROM
(SELECT DISTINCT actors.id, actors.first_name, actors.last_name 
FROM actors INNER JOIN movie_actors ON actors.id=movie_actors.actor_id WHERE movie_actors.movie_id=movies.id) AS t2) AS actors
FROM movies WHERE id = ?`
	if err := s.db.QueryRow(query, id).Scan(&movie.Id, &movie.Name, &movie.Genres, &movie.Actors); err != nil {
		return types.MovieDetail{}, err
	}

	return movie, nil
}
