-- add case insensitive text for email
CREATE EXTENSION IF NOT EXISTS citext;

-- create users table
CREATE TABLE IF NOT EXISTS users (
  id        bigserial PRIMARY KEY,
	email     citext UNIQUE NOT NULL,
	username  text UNIQUE NOT NULL,
	password  bytea NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW()
);
