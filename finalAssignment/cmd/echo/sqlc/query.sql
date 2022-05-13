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

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY id DESC;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

-- name: CreateList :execresult
INSERT INTO lists (
  name, userId
) VALUES (
?, ?
);

-- name: ListLists :many
SELECT * FROM lists
ORDER BY id DESC;

-- name: ListListsByUserId :many
SELECT * FROM lists
WHERE userId = ?;

-- name: DeleteLists :exec
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

-- name: GetUserByDate :one
SELECT * FROM users
WHERE datestamp = CURRENT_TIMESTAMP;

-- name: UpdateUsers :execresult
UPDATE users
SET datestamp = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE username = ?;

