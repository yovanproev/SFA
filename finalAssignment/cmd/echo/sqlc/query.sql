-- query.sql

-- name: CreateTask :execresult
INSERT INTO tasks (
  text, listId, userId, completed
) VALUES (
 ?, ?, ?, ?
);

-- name: UpdateTask :execresult
UPDATE tasks
SET completed = true
WHERE id = ?;

-- name: ListTasksByUserId :many
SELECT * FROM tasks
WHERE userId = ?;

-- name: DeleteTaskById :exec
DELETE FROM tasks
WHERE id = ?;

-- name: CreateList :execresult
INSERT INTO lists (
  name, userId
) VALUES (
?, ?
);

-- name: ListListsByUserId :many
SELECT * FROM lists
WHERE userId = ?;

-- name: DeleteListsById :exec
DELETE FROM lists
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
  username, password, datestamp
) VALUES (
 ?, ?, ?
);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

-- name: DeleteUserByUsername :exec
DELETE FROM users
WHERE username = ?;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = ?;
