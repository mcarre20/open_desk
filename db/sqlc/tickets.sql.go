// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tickets.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTicket = `-- name: CreateTicket :one
Insert Into tickets(
    user_id,
    description
)Values(
    $1,$2
)
Returning id, user_id, assigned_to, description, status, priority, category_id, updated_at, created_at
`

type CreateTicketParams struct {
	UserID      uuid.UUID `json:"user_id"`
	Description string    `json:"description"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, createTicket, arg.UserID, arg.Description)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AssignedTo,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.CategoryID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getTicket = `-- name: GetTicket :one
Select id, user_id, assigned_to, description, status, priority, category_id, updated_at, created_at From tickets
Where id = $1 LIMIT 1
`

func (q *Queries) GetTicket(ctx context.Context, id int64) (Ticket, error) {
	row := q.db.QueryRow(ctx, getTicket, id)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AssignedTo,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.CategoryID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getTicketList = `-- name: GetTicketList :many
Select id, user_id, assigned_to, description, status, priority, category_id, updated_at, created_at From tickets
Limit $1
Offset $2
`

type GetTicketListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetTicketList(ctx context.Context, arg GetTicketListParams) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, getTicketList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticket{}
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.AssignedTo,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.CategoryID,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTicket = `-- name: UpdateTicket :one
Update tickets 
Set
    assigned_to = $2,
    status = $3,
    priority = $4,
    category_id = $5,
    updated_at = Now()
Where id = $1
Returning id, user_id, assigned_to, description, status, priority, category_id, updated_at, created_at
`

type UpdateTicketParams struct {
	ID         int64     `json:"id"`
	AssignedTo uuid.UUID `json:"assigned_to"`
	Status     int32     `json:"status"`
	Priority   int32     `json:"priority"`
	CategoryID uuid.UUID `json:"category_id"`
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, updateTicket,
		arg.ID,
		arg.AssignedTo,
		arg.Status,
		arg.Priority,
		arg.CategoryID,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AssignedTo,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.CategoryID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
