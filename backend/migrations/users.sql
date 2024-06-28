CREATE IF NO EXIST SCHEMA users;

CREATE TABLE IF NOT EXISTS users.users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users.users (email);