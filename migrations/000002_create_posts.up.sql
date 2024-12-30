-- create posts table
CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  title text NOT NULL,
  user_id bigint NOT NULL,
  content text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

-- fk constraint for user_id
ALTER TABLE posts
ADD CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id);
