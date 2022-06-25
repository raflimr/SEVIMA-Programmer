// Code generated by sqlc. DO NOT EDIT.
// source: riwayat_penjualan.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteRiwayatPenjualan = `-- name: DeleteRiwayatPenjualan :exec
DELETE FROM riwayat_penjualan WHERE id = ?
`

func (q *Queries) DeleteRiwayatPenjualan(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRiwayatPenjualan, id)
	return err
}

const getAllRiwayatPenjualan = `-- name: GetAllRiwayatPenjualan :many
SELECT
   u.username,
   p.id, p.user_id, p.berat_sampah, p.jenis_sampah, p.harga_sampah, p.status, p.keuntungan, p.created_at
FROM
  riwayat_penjualan rj
JOIN
  user u ON rj.user_id = u.id
JOIN
  penjualan p ON rj.penjualan_id = p.id
`

type GetAllRiwayatPenjualanRow struct {
	Username    string        `json:"username"`
	ID          int32         `json:"id"`
	UserID      sql.NullInt32 `json:"user_id"`
	BeratSampah int32         `json:"berat_sampah"`
	JenisSampah string        `json:"jenis_sampah"`
	HargaSampah int32         `json:"harga_sampah"`
	Status      string        `json:"status"`
	Keuntungan  int32         `json:"keuntungan"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (q *Queries) GetAllRiwayatPenjualan(ctx context.Context) ([]GetAllRiwayatPenjualanRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllRiwayatPenjualan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllRiwayatPenjualanRow
	for rows.Next() {
		var i GetAllRiwayatPenjualanRow
		if err := rows.Scan(
			&i.Username,
			&i.ID,
			&i.UserID,
			&i.BeratSampah,
			&i.JenisSampah,
			&i.HargaSampah,
			&i.Status,
			&i.Keuntungan,
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

const getByIDRiwayatPenjualan = `-- name: GetByIDRiwayatPenjualan :many
SELECT
   u.username,
   p.id, p.user_id, p.berat_sampah, p.jenis_sampah, p.harga_sampah, p.status, p.keuntungan, p.created_at
FROM
  riwayat_penjualan rj
JOIN
  user u ON rj.user_id = u.id
JOIN
  penjualan p ON rj.penjualan_id = p.id
WHERE
  u.id = ?
`

type GetByIDRiwayatPenjualanRow struct {
	Username    string        `json:"username"`
	ID          int32         `json:"id"`
	UserID      sql.NullInt32 `json:"user_id"`
	BeratSampah int32         `json:"berat_sampah"`
	JenisSampah string        `json:"jenis_sampah"`
	HargaSampah int32         `json:"harga_sampah"`
	Status      string        `json:"status"`
	Keuntungan  int32         `json:"keuntungan"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (q *Queries) GetByIDRiwayatPenjualan(ctx context.Context, id int32) ([]GetByIDRiwayatPenjualanRow, error) {
	rows, err := q.db.QueryContext(ctx, getByIDRiwayatPenjualan, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetByIDRiwayatPenjualanRow
	for rows.Next() {
		var i GetByIDRiwayatPenjualanRow
		if err := rows.Scan(
			&i.Username,
			&i.ID,
			&i.UserID,
			&i.BeratSampah,
			&i.JenisSampah,
			&i.HargaSampah,
			&i.Status,
			&i.Keuntungan,
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
