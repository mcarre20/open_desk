-- name: CreateUser :one
Insert Into users(
    username,
    first_name,
    last_name,
    email,
    hashed_password,
    user_role
)Values(
    $1,$2,$3,$4,$5,$6
)
Returning *;

-- name: GetUser :one
Select * From users
Where id = $1 Limit 1;

-- name: GetUserByUserName :one
Select * From users
Where username = $1 Limit 1;

-- name: GetUserList :many
Select * From users
Where active = TRUEgit
Order By username
Limit $1
OFFSET $2;


-- name: UpdateUserInfo :one
Update users 
Set
    first_name = $2,
    last_name = $3,
    user_role = $4,
    email =$5,
    updated_at = now()
Where id = $1
Returning *;

-- name: UpdateUserPassword :one
Update users
Set
    hashed_password = $2,
    password_updated_at = now()
Where id = $1
Returning *;

-- name: DeactivateUser :exec
Update users 
Set
    active = TRUE,
    updated_at = now()
Where id = $1;