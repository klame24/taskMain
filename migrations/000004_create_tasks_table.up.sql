CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    project_id INTEGER REFERENCES projects(id),
    title VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(1000),
    status_id INTEGER REFERENCES statuses(id),
    created_at TIMESTAMP NOT NULL
)