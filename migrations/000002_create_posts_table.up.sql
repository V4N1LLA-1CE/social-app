CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  content text NOT NULL,
  title text NOT NULL,
  user_id bigint NOT NULL,
  tags text[] NOT NULL DEFAULT '{}',
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE posts
ADD CONSTRAINT users_posts_fk FOREIGN KEY (user_id)
  REFERENCES users(id) ON DELETE CASCADE;
