-- +goose Up

CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  url TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP,
  feed_id UUID NOT NULL,
  UNIQUE(feed_id, url),
  FOREIGN KEY(feed_id)
  REFERENCES feeds(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_posts_published_at ON posts (published_at);

-- +goose Down

DROP INDEX IF EXISTS idx_posts_published_at;

DROP TABLE posts;
