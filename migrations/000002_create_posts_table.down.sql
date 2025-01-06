-- remove fk
ALTER TABLE posts
DROP CONSTRAINT fk_users_posts;

-- then drop the table
DROP TABLE IF EXISTS posts;
