CREATE TABLE
    IF NOT EXISTS fpd.transactions (
        id UUID PRIMARY KEY,
        user_id VARCHAR(255) NOT NULL,
        bond_id UUID NOT NULL,
        status VARCHAR(20) NOT NULL,
        quantity INTEGER NOT NULL,
        price DECIMAL(14, 4) NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL,
        deleted_at INTEGER,
        FOREIGN KEY (user_id) REFERENCES users.users (id),
        FOREIGN KEY (bond_id) REFERENCES fpd.bonds (id)
    );

CREATE INDEX IF NOT EXISTS transactions_user_id_idx ON fpd.transactions (user_id);