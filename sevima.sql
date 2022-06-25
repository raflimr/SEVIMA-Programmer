CREATE TABLE `pembelian` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `nama_barang` varchar(255)  NOT NULL,
  `harga_barang` int(255)  NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT "pengecekan",
  `created_at` datetime NOT NULL
);

CREATE TABLE `penjualan` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `berat_sampah` int(11)  NOT NULL,
  `jenis_sampah` varchar(255)  NOT NULL,
  `harga_sampah` int(11)  NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT "pengecekan",
  `keuntungan` int(11)  NOT NULL,
  `created_at` datetime NOT NULL
);

CREATE TABLE `riwayat_pembelian` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `pembelian_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL
);

CREATE TABLE `riwayat_penjualan` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `penjualan_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL
);

CREATE TABLE `user` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `photo` varchar(255) DEFAULT NULL,
  `username` varchar(255)  NOT NULL,
  `password` varchar(255)  NOT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `saldo` int(255) NOT NULL DEFAULT 0,
  `total_sampah` int(255) NOT NULL DEFAULT 0,
  `role` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL
);

ALTER TABLE `penjualan` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `pembelian` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `riwayat_pembelian` ADD FOREIGN KEY (`pembelian_id`) REFERENCES `pembelian` (`id`);

ALTER TABLE `riwayat_pembelian` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `riwayat_penjualan` ADD FOREIGN KEY (`penjualan_id`) REFERENCES `penjualan` (`id`);

ALTER TABLE `riwayat_penjualan` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
