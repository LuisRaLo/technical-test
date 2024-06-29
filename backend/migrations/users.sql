CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE
    IF NOT EXISTS users.users (
        id varchar(255) PRIMARY KEY NOT NULL,
        email VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL,
        deleted_at INTEGER
    );

CREATE INDEX IF NOT EXISTS users_email_idx ON users.users (email);