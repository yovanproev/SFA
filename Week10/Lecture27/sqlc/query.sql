-- query.sql

-- name: CreateStory :execresult
INSERT INTO items (
  Title, Score, DateStamp
) VALUES (
 ?, ?, ?
);

-- name: GetLastStory :one
SELECT * FROM items
ORDER BY DateStamp DESC
LIMIT 1;

-- name: ListStories :many
SELECT * FROM items
ORDER BY DateStamp DESC
LIMIT 10;


-- name: DeleteStory :exec
DELETE FROM items
WHERE id = ?;