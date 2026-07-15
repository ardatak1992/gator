-- +goose Up
CREATE TABLE
  feeds_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    UNIQUE (feed_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (feed_id) REFERENCES feeds (id)
  );


-- +goose Down
DROP TABLE feeds_follows;
