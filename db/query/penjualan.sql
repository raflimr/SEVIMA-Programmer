-- name: CreatePenjualan :exec
INSERT INTO penjualan (
  berat_sampah, jenis_sampah, harga_sampah, keuntungan
) VALUES (
  ?, ?,? ,?
);

-- name: UpdateStatusPenjualan :exec
UPDATE penjualan SET status = "success" WHERE user_id = ?;