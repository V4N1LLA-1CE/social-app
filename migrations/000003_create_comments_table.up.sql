CREATE TABLE IF NOT EXISTS comments (
  id bigserial PRIMARY KEY,
  post_id bigserial NOT NULL,
  user_id bigserial NOT NULL,
  content text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE comments
   add constraint comments_user_fk FOREIGN KEY (user_id)
     REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE comments
   ADD CONSTRAINT comments_posts_fk FOREIGN KEY(post_id)
     REFERENCES posts(id) ON DELETE CASCADE;
