-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: ExistUser :one
SELECT EXISTS(select * from users);