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

INSERT INTO
    users.users (
        id,
        email,
        name,
        created_at,
        updated_at,
        deleted_at
    )
VALUES
    (
        'h5kj5hdrAUPdYtoq2faqcekz8ih1',
        'technical.challenge.user@yopmail.com',
        'Technical Challenge User',
        1719689016,
        1719689016,
        NULL
    ),
    (
        'dzvXQ3xlsGTsHYgtoyrMDsFuYnh2',
        'technical.challenge.user1@yopmail.com',
        'User Test 2',
        1719692401,
        1719692401,
        0
    );