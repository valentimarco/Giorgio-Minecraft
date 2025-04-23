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


-- name: CreateServer :one
INSERT INTO servers(
  type,name,memory,max_player,aika_flag
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetServerByID :one
SELECT * FROM servers
WHERE id = $1;