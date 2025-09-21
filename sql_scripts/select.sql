-- Created this file to try different queries.
-- First I test queries here then write in Go.

-- Select movie for movie list page(add pagination later, add poster later)
SELECT id, name, release_date, duration_in_minutes, rating FROM movies;

-- Select movie for movie detail page
SELECT movies.id, movies.name, 
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'name', name)) 
FROM
(SELECT DISTINCT genres.id, genres.name
FROM genres INNER JOIN movie_genres ON genres.id=movie_genres.genre_id WHERE movie_genres.movie_id=movies.id) AS t1) AS genres,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('id', id, 'first_name', first_name, 'last_name', last_name)) 
FROM
(SELECT DISTINCT actors.id, actors.first_name, actors.last_name 
FROM actors INNER JOIN movie_actors ON actors.id=movie_actors.actor_id WHERE movie_actors.movie_id=movies.id) AS t2) AS actors,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('name', name, 'type', type))
FROM
(SELECT DISTINCT movie_medias.name, movie_medias.type 
FROM movie_medias WHERE movie_medias.movie_id=movies.id) AS t3) AS medias
FROM movies;

-- Select ratings
SELECT m.id, m.movie_id, m.rating, m.comment, p.first_name, p.last_name
FROM movie_ratings m 
INNER JOIN profiles p ON p.user_id=m.user_id 
WHERE m.movie_id = 1;

-- Select ratings of the user
SELECT r.id, r.movie_id, r.rating, r.comment, m.name 
FROM movie_ratings r 
INNER JOIN movies m ON m.id=r.movie_id 
WHERE r.user_id = '37890809-bc1d-4266-b605-c4e83c2a6a62';

SELECT EXISTS(SELECT 1 FROM movies WHERE id=1) AS is_exist;

SELECT movies.id, movies.name, 
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
FROM movies WHERE id = 1