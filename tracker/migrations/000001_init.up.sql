CREATE TABLE users(
        user_id SERIAL PRIMARY KEY,
        passport_serie VARCHAR(4) NOT NULL,
        passport_number VARCHAR(6) NOT NULL
);
