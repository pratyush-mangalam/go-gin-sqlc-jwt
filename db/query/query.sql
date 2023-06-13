-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetAllUsers :many
select * from users;

-- name: CreateUser :one
insert into users (name, email, number, password)
values ($1, $2, $3, $4) returning *;