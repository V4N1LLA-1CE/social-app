-- drop the join table first to handle foreign key dependencies
DROP TABLE IF EXISTS users_permissions;

-- drop the main tables
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS users;

-- drop citext extension
DROP EXTENSION IF EXISTS citext;
