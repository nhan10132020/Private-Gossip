CREATE TABLE users (
    "id" bigserial PRIMARY KEY,
    "username" text NOT NULL,
    "email" citext UNIQUE NOT NULL,
    "password_hash" bytea NOT NULL
);