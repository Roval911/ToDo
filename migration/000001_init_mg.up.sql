-- Create users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(150) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);

-- Create tasks table
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       completed BOOLEAN DEFAULT FALSE,
                       user_id INTEGER NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);