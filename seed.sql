CREATE TABLE users_md (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cellphone VARCHAR(20) NOT NULL,
    password VARCHAR(255) NOT NULL
);

