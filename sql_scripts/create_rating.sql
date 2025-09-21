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