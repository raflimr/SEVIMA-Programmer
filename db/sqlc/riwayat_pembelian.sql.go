// Code generated by sqlc. DO NOT EDIT.
// source: riwayat_pembelian.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteRiwayatPembelian = `-- name: DeleteRiwayatPembelian :exec
DELETE FROM riwayat_pembelian WHERE id = ?
`

func (q *Queries) DeleteRiwayatPembelian(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRiwayatPembelian, id)
	return err
}

const getAllRiwayatPembelian = `-- name: GetAllRiwayatPembelian :many
SELECT
   u.username,
   p.id, p.user_id, p.nama_barang, p.harga_barang, p.status, p.created_at
FROM
  riwayat_pembelian rb
JOIN
  user u ON rb.user_id = u.id
JOIN
  pembelian p ON rb.pembelian_id = p.id
`

type GetAllRiwayatPembelianRow struct {
	Username    string        `json:"username"`
	ID          int32         `json:"id"`
	UserID      sql.NullInt32 `json:"user_id"`
	NamaBarang  string        `json:"nama_barang"`
	HargaBarang int32         `json:"harga_barang"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (q *Queries) GetAllRiwayatPembelian(ctx context.Context) ([]GetAllRiwayatPembelianRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllRiwayatPembelian)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllRiwayatPembelianRow
	for rows.Next() {
		var i GetAllRiwayatPembelianRow
		if err := rows.Scan(
			&i.Username,
			&i.ID,
			&i.UserID,
			&i.NamaBarang,
			&i.HargaBarang,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getByIDRiwayatPembelian = `-- name: GetByIDRiwayatPembelian :many
SELECT
   u.username,
   p.id, p.user_id, p.nama_barang, p.harga_barang, p.status, p.created_at
FROM
  riwayat_pembelian rb
JOIN
  user u ON rb.user_id = u.id
JOIN
  pembelian p ON rb.pembelian_id = p.id
WHERE
  u.id = ?
`

type GetByIDRiwayatPembelianRow struct {
	Username    string        `json:"username"`
	ID          int32         `json:"id"`
	UserID      sql.NullInt32 `json:"user_id"`
	NamaBarang  string        `json:"nama_barang"`
	HargaBarang int32         `json:"harga_barang"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (q *Queries) GetByIDRiwayatPembelian(ctx context.Context, id int32) ([]GetByIDRiwayatPembelianRow, error) {
	rows, err := q.db.QueryContext(ctx, getByIDRiwayatPembelian, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetByIDRiwayatPembelianRow
	for rows.Next() {
		var i GetByIDRiwayatPembelianRow
		if err := rows.Scan(
			&i.Username,
			&i.ID,
			&i.UserID,
			&i.NamaBarang,
			&i.HargaBarang,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
