CREATE TABLE IF NOT EXISTS users
(
    id UUID NOT NULL PRIMARY KEY,
    login VARCHAR(255)  NOT NULL,
    password VARCHAR(255)  NOT NULL
);
