-- Creates all the tables

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

CREATE TABLE actors (
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE genres (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE movie_genres (
    id INT AUTO_INCREMENT NOT NULL,
    movie_id INT NOT NULL,
    genre_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    FOREIGN KEY(genre_id) REFERENCES genres(id),
    PRIMARY KEY(id)
);

CREATE TABLE movie_actors (
    id INT AUTO_INCREMENT NOT NULL,
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    FOREIGN KEY(actor_id) REFERENCES actors(id),
    PRIMARY KEY(id)
);

CREATE TABLE users (
    id VARCHAR(36) NOT NULL,
    email VARCHAR(254) NOT NULL,
    password VARCHAR(60) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE profiles (
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    about_me VARCHAR(255) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    PRIMARY KEY(id)
);

CREATE TABLE movie_ratings (
    id INT AUTO_INCREMENT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    movie_id INT NOT NULL,
    rating FLOAT NOT NULL,
    comment VARCHAR(255) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    PRIMARY KEY(id)
);

CREATE TABLE movie_medias (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    type TINYINT NOT NULL,
    movie_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES movies(id),
    PRIMARY KEY(id)
);