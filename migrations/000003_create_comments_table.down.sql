ALTER TABLE comments
DROP CONSTRAINT comments_user_fk;

ALTER TABLE comments
   DROP CONSTRAINT comments_posts_fk;

DROP TABLE IF EXISTS comments;
