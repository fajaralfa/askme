CREATE DATABASE IF NOT EXISTS askme;
USE askme;
CREATE TABLE users (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(60) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    photo VARCHAR(100) NOT NULL
);

CREATE TABLE questions (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    question VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    seen BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);