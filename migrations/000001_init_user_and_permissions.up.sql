-- Enable citext extension
CREATE EXTENSION IF NOT EXISTS citext;

-- create users table
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  email citext UNIQUE NOT NULL,
  username text UNIQUE NOT NULL,
  password bytea NOT NULL,
  activated bool NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

-- create permissions table
CREATE TABLE IF NOT EXISTS permissions (
  id bigserial PRIMARY KEY,
  scope text NOT NULL,
  description text NOT NULL
);

-- create join table between users and permissions
CREATE TABLE IF NOT EXISTS users_permissions (
  user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
  permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
  PRIMARY KEY (user_id, permission_id)
);
