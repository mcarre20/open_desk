-- name: CreateComment :one
Insert Into comments(
    user_id,
    ticket_id,
    comments,
    customer_visible
)Values(
    $1,$2,$3,$4
)
Returning *;
-- name: GetComment :one
Select * From comments
Where id = $1
Limit 1;

-- name: GetTicketComments :many
Select * From comments
Where ticket_id = $1
ORDER BY created_at DESC;

-- name: UpdateComment :one
Update comments
Set
    comments = $2,
    customer_visible = $3,
    updated_at = NOW()
Where id = $1
Returning *;

