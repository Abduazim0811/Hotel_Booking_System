CREATE TABLE IF NOT EXISTS users(
    userid serial PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    age int NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);