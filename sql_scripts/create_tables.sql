DROP TABLE IF EXISTS movies;

CREATE TABLE movies (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NOT NULL,
    rating FLOAT NOT NULL,
    duration_in_minutes INT NOT NULL,
    release_date DATETIME NOT NULL,
    language_code VARCHAR(2) NOT NULL,
    number_of_ratings INT NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO movies 
(name, description, rating, duration_in_minutes, release_date, language_code, number_of_ratings) 
VALUES 
('Madame Web', 
'Forced to confront her past, Cassandra Webb, a Manhattan paramedic that may have clairvoyant abilities, forms a relationship with three young women destined for powerful futures, if they can survive their threatening present.',
4.1, 116, '2024-02-14', 'en', 107000),
('Batman Begins', 
'After witnessing his parents death, billionaire Bruce Wayne learns the art of fighting to confront injustice. When he returns to Gotham as Batman, he must stop a secret society that intends to destroy the city.',
8.2, 140, '2005-06-10', 'en', 1700000),
('Patch Adams', 
'The true story of the heroic Hunter "Patch" Adams whos determined to become a medical doctor because of his desire to help other people. He ventures where no doctor ever ventured before utilizing humor and pathos.',
6.9, 115, '1998-12-25', 'en', 128000),
('Angel-A', 
'A beautiful woman helps an inept scam artist get his game together.',
7.0, 91, '2005-12-21', 'fr', 37000);

DROP TABLE IF EXISTS actors;

CREATE TABLE actors (
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO actors 
(first_name, last_name) 
VALUES 
('Dakota', 'Johnson'), ('Sydney', 'Sweeney'), ('Isabela', 'Merced'),
('Christian', 'Bale'), ('Michael', 'Caine'), ('Ken', 'Watanabe'),
('Robin', 'Williams'), ('Daniel', 'London'), ('Monica', 'Potter'),
('Rie', 'Rasmussen'), ('Jamel', 'Debbouze'), ('Gilbert', 'Melki');

DROP TABLE IF EXISTS genres;

CREATE TABLE genres (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO genres (name) 
VALUES 
('Superhero'), ('Action'), ('Thriller'), ('Action Epic'), ('Epic'), ('Tragedy'), ('Crime'), ('Drama'),
('Docudrama'), ('Medical Drama'), ('Biography'), ('Comedy'), ('Romance'), ('Fantasy');

DROP TABLE IF EXISTS movie_genres;

CREATE TABLE movie_genres (
    id INT AUTO_INCREMENT NOT NULL,
    movie_id INT NOT NULL,
    genre_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    FOREIGN KEY(genre_id) REFERENCES genres(id),
    PRIMARY KEY(id)
);

INSERT INTO movie_genres 
(movie_id, genre_id) 
VALUES
(1, 1), (1, 2), (1, 3),
(2, 4), (2, 5), (2, 1), (2, 6), (2, 7), (2, 8), (2, 3),
(3, 9), (3, 10), (3, 11), (3, 12), (3, 8), (3, 13),
(4, 12), (4, 8), (4, 14), (4, 13);

DROP TABLE IF EXISTS movie_actors;

CREATE TABLE movie_actors (
    id INT AUTO_INCREMENT NOT NULL,
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    FOREIGN KEY(actor_id) REFERENCES actors(id),
    PRIMARY KEY(id)
);

INSERT INTO movie_actors 
(movie_id, actor_id) 
VALUES 
(1, 1), (1, 2), (1, 3), (2, 4), (2, 5), (2, 6), (3, 7), (3, 8), (3, 9), (4, 10), (4, 11), (4, 12);