-- name: CreateUser :one
INSERT INTO users (username, password, phone, first_name, last_name, is_active, birth_date, role, address ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning guid;

-- name: GetUser :one
select * from users where guid = $1;

-- name: GetAllUsers :many
select * from users limit $1 offset $2;

-- name: UpdateUserPassword :exec
update users set password = $3 where guid = $1 and password = $2;

-- name: DeleteUser :exec
update users set deleted_at = now() where guid = $1;    

-- name: LoginUser :one
select guid, username, password, role, is_active from users where username = $1;