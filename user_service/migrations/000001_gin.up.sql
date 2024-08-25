CREATE TABLE IF NOT EXISTS users(
    userid serial PRIMARY KEY,
    username VARCHAR(50),
    age int,
    email VARCHAR(50),
    password VARCHAR(100)
);