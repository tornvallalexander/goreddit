-- name: CreateSubreddit :one
INSERT INTO subreddits (
    name,
    moderator,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetSubreddit :one
SELECT * FROM subreddits
WHERE name = $1
LIMIT 1;

-- name: DeleteSubreddit :one
DELETE FROM subreddits
WHERE name = $1
RETURNING *;