/* CREATE SCHEMA FIRST PARTY DATA */
CREATE IF NO EXIST SCHEMA fpd;

CREATE TABLE IF NOT EXISTS fpd.bonds (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(14, 4) NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    deleted_at INTEGER,
    FOREIGN KEY (user_id) REFERENCES users.users(id)
);

CREATE INDEX IF NOT EXISTS bonds_user_id_idx ON fpd.bonds (user_id);