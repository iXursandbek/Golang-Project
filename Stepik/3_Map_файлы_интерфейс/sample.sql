CREATE DATABASE learning;


CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    email VARCHAR(60) NOT NULL,
    birthday DATE NOT NULL
);
 

INSERT INTO users (
    id,
    name,
    email,
    birthday)
VALUES ('Xursandbek','ixursand@mail.ru', DATE '1995-10-14');