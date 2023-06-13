CREATE TABLE USERS
(
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    email varchar(100) UNIQUE NOT NULL,
    number varchar(10) UNIQUE NOT NULL,
    password varchar(100) NOT NULL
);