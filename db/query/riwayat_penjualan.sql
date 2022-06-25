-- name: GetAllRiwayatPenjualan :many
SELECT
   u.username,
   p.*
FROM
  riwayat_penjualan rj
JOIN
  user u ON rj.user_id = u.id
JOIN
  penjualan p ON rj.penjualan_id = p.id;


-- name: GetByIDRiwayatPenjualan :many
SELECT
   u.username,
   p.*
FROM
  riwayat_penjualan rj
JOIN
  user u ON rj.user_id = u.id
JOIN
  penjualan p ON rj.penjualan_id = p.id
WHERE
  u.id = ?;

-- name: DeleteRiwayatPenjualan :exec
DELETE FROM riwayat_penjualan WHERE id = ?;