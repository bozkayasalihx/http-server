-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;


-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAuthor :exec
DELETE FROM accounts 
WHERE id = $1;

-- name: UpdateAuthor :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;
