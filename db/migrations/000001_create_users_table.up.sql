-- Enable the uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the users table
CREATE TABLE users
(
    user_id     UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) UNIQUE,
    password    VARCHAR(255),
    roles      VARCHAR(255)[],
    created_at  timestamptz NOT NULL DEFAULT (now()),
    updated_at  timestamptz NOT NULL DEFAULT (now())
);

-- Create an index on the name column for better performance
CREATE INDEX idx_users_name ON users (name);