-- name: CreatePost :one
INSERT INTO posts (
    "user",
    title,
    content,
    "subreddit"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
WHERE subreddit = $1
ORDER BY id
LIMIT $2;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts
SET title = $1, content = $2
WHERE id = $3
RETURNING *;