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