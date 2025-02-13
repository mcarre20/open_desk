-- name: CreateCategory :one
Insert Into caterogies(
    category
)Values(
    $1
)
Returning *;

-- name: GetAllCategories :many
Select * From caterogies;

-- name: UpdateCategory :one
Update caterogies
Set
    category = $2
Where id = $1
Returning *;

