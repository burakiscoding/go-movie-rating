-- Created this file to try different queries.
-- First I tested queries here.

-- Select movie for movie list page(add pagination later, add poster later)
SELECT id, name, release_date, duration_in_minutes, rating FROM movies;

-- Select movie for movie detail page
SELECT movies.id, movies.name, 
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'name', name)) 
FROM
(SELECT DISTINCT genres.id, genres.name
FROM genres INNER JOIN movie_genres ON genres.id=movie_genres.genre_id WHERE movie_genres.movie_id=movies.id) AS t1) AS genres,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'first_name', first_name,'last_name',  last_name)) 
FROM
(SELECT DISTINCT actors.id, actors.first_name, actors.last_name 
FROM actors INNER JOIN movie_actors ON actors.id=movie_actors.actor_id WHERE movie_actors.movie_id=movies.id) AS t2) AS actors
FROM movies;