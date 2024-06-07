-- name: GetUserByUsername :one
SELECT au.*
FROM "auth_user" au
WHERE au.username = $1
LIMIT 1;

-- name: GetUserById :one
SELECT au.id,
       first_name,
       last_name,
       username,
       email,
       phone,
       '' as password,
       is_superuser,
       is_active,
       is_staff,
       is_bot,
       last_login,
       created_at,
       updated_at
FROM "auth_user" au
WHERE au.id = $1
LIMIT 1;

-- name: SaveAuthTokenPair :exec
INSERT INTO auth_user_token (user_id, token, token_type, expires_at, created_at)
VALUES ($1, $2, 'access', $3, $4),
       ($1, $5, 'refresh', $3, $4);

-- name: GetActiveRefreshToken :one
SELECT aut.*
FROM auth_user_token aut
WHERE aut.token = $1
  AND aut.token_type = 'refresh'
  AND aut.is_active = TRUE
LIMIT 1;
