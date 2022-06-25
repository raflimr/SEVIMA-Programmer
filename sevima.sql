CREATE TABLE `pembelian` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `nama_barang` int(11) DEFAULT NULL,
  `harga_barang` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT "pengecekan",
  `created_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `penjualan` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `berat_sampah` int(11) DEFAULT NULL,
  `jenis_sampah` varchar(255) DEFAULT NULL,
  `harga_sampah` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT "pengecekan",
  `keuntungan` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `riwayat` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `pembelian_id` int(11) DEFAULT NULL,
  `penjualan_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `user` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `photo` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `saldo` int(255) DEFAULT 0,
  `total_sampah` int(255) DEFAULT 0,
  `role` varchar(255) DEFAULT "user",
  `created_at` datetime NOT NULL DEFAULT (now())
);

ALTER TABLE `penjualan` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `pembelian` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `riwayat` ADD FOREIGN KEY (`penjualan_id`) REFERENCES `penjualan` (`id`);

ALTER TABLE `riwayat` ADD FOREIGN KEY (`pembelian_id`) REFERENCES `pembelian` (`id`);

ALTER TABLE `riwayat` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
