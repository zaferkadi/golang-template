-- name: CreateGenre :one
INSERT INTO genres
    (
    name, description
    )
VALUES
    (
        $1, $2
)
RETURNING *;

-- name: GetGenre :one
SELECT *
FROM genres
WHERE id = $1
LIMIT 1;

-- name: ListGenres :many
SELECT *
FROM genres
ORDER BY id
LIMIT $1
OFFSET
$2;

-- name: UpdateGenre :one
UPDATE genres
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1;