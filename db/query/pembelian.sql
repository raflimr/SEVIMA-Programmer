-- name: CreatePembelian :exec
INSERT INTO pembelian (
  nama_barang, harga_barang
) VALUES (
  ?, ?
);

-- name: UpdateStatusPembelian :exec
UPDATE pembelian SET status = ? WHERE user_id = ?;