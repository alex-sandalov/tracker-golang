CREATE TABLE users(
        user_id SERIAL PRIMARY KEY,
        passport_serie VARCHAR(4) NOT NULL,
        passport_number VARCHAR(6) NOT NULL
);

CREATE TABLE user_tasks(
        task_id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        description VARCHAR(255) NOT NULL,
        start_time TIMESTAMP NOT NULL,
        end_time TIMESTAMP,
        active BOOLEAN NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(user_id)
);
