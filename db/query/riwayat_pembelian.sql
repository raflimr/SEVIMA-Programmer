-- name: GetAllRiwayatPembelian :many
SELECT
   u.username,
   p.*
FROM
  riwayat_pembelian rb
JOIN
  user u ON rb.user_id = u.id
JOIN
  pembelian p ON rb.pembelian_id = p.id;


-- name: GetByIDRiwayatPembelian :many
SELECT
   u.username,
   p.*
FROM
  riwayat_pembelian rb
JOIN
  user u ON rb.user_id = u.id
JOIN
  pembelian p ON rb.pembelian_id = p.id
WHERE
  u.id = ?;

-- name: DeleteRiwayatPembelian :exec
DELETE FROM riwayat_pembelian WHERE id = ?;