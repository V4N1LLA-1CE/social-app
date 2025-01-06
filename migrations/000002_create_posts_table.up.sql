-- create posts table
CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  content text NOT NULL,
  title text NOT NULL,
  user_id bigint NOT NULL,
  tags []text NOT NULL DEFAULT '{}',
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

-- add fk constraint
ALTER TABLE posts
ADD CONSTRAINT fk_users_posts FOREIGN KEY (user_id)
  REFERENCES users(id)
