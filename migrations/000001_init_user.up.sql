-- Enable citext extension
CREATE EXTENSION IF NOT EXISTS citext;

-- create users table
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  email citext UNIQUE NOT NULL,
  password bytea NOT NULL,
  activated bool NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);
