-- name: GetUserByEmail :one
SELECT id, email, password_hash, role, first_name, last_name, created_at, updated_at
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash, role, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, password_hash, role, first_name, last_name, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, email, password_hash, role, first_name, last_name, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetDoctors :many
SELECT id, email, role, first_name, last_name, created_at, updated_at
FROM users
WHERE role = 'doctor'
ORDER BY first_name, last_name;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg(email), email),
    password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
    role = COALESCE(sqlc.narg(role), role),
    first_name = COALESCE(sqlc.narg(first_name), first_name),
    last_name = COALESCE(sqlc.narg(last_name), last_name),
    updated_at = NOW()
WHERE id = $1
RETURNING id, email, password_hash, role, first_name, last_name, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
