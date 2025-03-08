// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
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
Returning id, username, first_name, last_name, email, hashed_password, user_role, active, created_at, updated_at, password_updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	UserRole       int32  `json:"user_role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.HashedPassword,
		arg.UserRole,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.UserRole,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PasswordUpdatedAt,
	)
	return i, err
}

const deactivateUser = `-- name: DeactivateUser :exec
Update users 
Set
    active = TRUE,
    updated_at = now()
Where id = $1
`

func (q *Queries) DeactivateUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deactivateUser, id)
	return err
}

const getUser = `-- name: GetUser :one
Select id, username, first_name, last_name, email, hashed_password, user_role, active, created_at, updated_at, password_updated_at From users
Where id = $1 Limit 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.UserRole,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PasswordUpdatedAt,
	)
	return i, err
}

const getUserList = `-- name: GetUserList :many
Select id, username, first_name, last_name, email, hashed_password, user_role, active, created_at, updated_at, password_updated_at From users
Where active = TRUEgit
Order By username
Limit $1
OFFSET $2
`

type GetUserListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUserList(ctx context.Context, arg GetUserListParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUserList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.HashedPassword,
			&i.UserRole,
			&i.Active,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PasswordUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserInfo = `-- name: UpdateUserInfo :one
Update users 
Set
    first_name = $2,
    last_name = $3,
    user_role = $4,
    email =$5,
    updated_at = now()
Where id = $1
Returning id, username, first_name, last_name, email, hashed_password, user_role, active, created_at, updated_at, password_updated_at
`

type UpdateUserInfoParams struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserRole  int32     `json:"user_role"`
	Email     string    `json:"email"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserInfo,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.UserRole,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.UserRole,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PasswordUpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
Update users
Set
    hashed_password = $2,
    password_updated_at = now()
Where id = $1
Returning id, username, first_name, last_name, email, hashed_password, user_role, active, created_at, updated_at, password_updated_at
`

type UpdateUserPasswordParams struct {
	ID             uuid.UUID `json:"id"`
	HashedPassword string    `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.UserRole,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PasswordUpdatedAt,
	)
	return i, err
}
