CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  username text UNIQUE NOT NULL,
  email citext UNIQUE NOT NULL,
  password bytea NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);
