CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE USER_ROLE AS ENUM ('admin', 'user');
CREATE TYPE USER_STATUS AS ENUM ('online', 'offline');
CREATE TYPE ACCOUNT_STATUS AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    login TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    status USER_STATUS NOT NULL DEFAULT 'online',
    role USER_ROLE NOT NULL DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
CREATE UNIQUE INDEX uq_users_login_lower ON users (lower(login));

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    currency TEXT NOT NULL,
    balance NUMERIC(19, 4) NOT NULL DEFAULT 0 CHECK (balance >= 0),
    status ACCOUNT_STATUS NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
CREATE INDEX idx_accounts_user_id ON accounts(user_id);