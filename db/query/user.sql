-- name: RegisterUser :exec
INSERT INTO user (
  username, password, role
) VALUES (
  ?, ?, ?
);

-- name: GetUserByID :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserByUsernameAndPassword :one
SELECT * FROM user WHERE username = ? AND password = ?;

-- name: GetAdminUser :one
SELECT * FROM user WHERE username = "admin";

-- name: UpdateSaldodanTotalSampah :exec
UPDATE user SET saldo = ?,  total_sampah = ? WHERE id = ?;

-- name: UpdateProfile :exec
UPDATE user SET photo = ?, password = ?, no_hp = ?,email = ? WHERE id = ?;