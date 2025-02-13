-- name: CreateTicket :one
Insert Into tickets(
    user_id,
    description
)Values(
    $1,$2
)
Returning *;

-- name: GetTicket :one
Select * From tickets
Where id = $1 LIMIT 1;

-- name: GetTicketList :many
Select * From tickets
Limit $1
Offset $2;

-- name: UpdateTicket :one
Update tickets 
Set
    assigned_to = $2,
    status = $3,
    priority = $4,
    category_id = $5,
    updated_at = Now()
Where id = $1
Returning *;
