CREATE TABLE users
(
    id       CHAR(36)    NOT NULL,
    username VARCHAR(20) NOT NULL UNIQUE,
    password CHAR(60)    NOT NULL,
    status   VARCHAR(30) NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX name ON users (username);