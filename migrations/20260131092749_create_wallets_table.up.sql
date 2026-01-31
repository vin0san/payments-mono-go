CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL UNIQUE,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_wallet_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);