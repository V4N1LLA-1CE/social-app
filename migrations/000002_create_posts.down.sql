-- drop foreign key constraint
ALTER TABLE posts
DROP CONSTRAINT IF EXISTS fk_user;

-- drop posts table
DROP TABLE IF EXISTS posts;
