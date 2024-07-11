CREATE TABLE users(
        id SERIAL PRIMARY KEY,
        passport_number INT NOT NULL,
        passport_serie INT NOT NULL,
        surname VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        patronymic VARCHAR(255),
        address VARCHAR(255) NOT NULL
);
