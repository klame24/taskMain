CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES users(id),
    title VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(1000),
    status_id INTEGER REFERENCES statuses(id),
    created_at TIMESTAMP NOT NULL
);