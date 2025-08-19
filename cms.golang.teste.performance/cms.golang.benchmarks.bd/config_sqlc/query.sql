-- name: GetById :one
SELECT id, name, created_at FROM "TbPerson"
WHERE id = $1 LIMIT 1;

-- name: GetAll :many
SELECT id, name, created_at FROM "TbPerson"
ORDER BY created_at ;

-- name: Create :one
INSERT INTO "TbPerson" (
  id, name, created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: Update :exec
UPDATE "TbPerson"
  set name = $2
WHERE id = $1;

-- name: Delete :exec
DELETE FROM "TbPerson"
WHERE id = $1;

-- name: Clear :exec
DELETE FROM "TbPerson"
WHERE name like 'SqlC%';