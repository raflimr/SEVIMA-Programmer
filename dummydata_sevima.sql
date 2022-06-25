-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Jun 2022 pada 22.56
-- Versi server: 10.4.20-MariaDB
-- Versi PHP: 7.4.22

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sevima`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `saldo` int(255) DEFAULT 0,
  `total_sampah` int(255) DEFAULT 0,
  `role` varchar(255) DEFAULT 'user',
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id`, `photo`, `username`, `password`, `no_hp`, `email`, `saldo`, `total_sampah`, `role`, `created_at`) VALUES
(1, NULL, 'rafli', 'ada', NULL, NULL, 1000, 100, 'user', '2022-06-25 14:11:43'),
(2, NULL, 'ADA', 'AJA', NULL, NULL, 0, 0, 'user', '2022-06-25 14:12:22'),
(3, NULL, 'username', 'password', NULL, NULL, 0, 0, 'user', '2022-06-25 16:36:30'),
(4, NULL, 'username', 'password', NULL, NULL, 0, 0, 'user', '2022-06-25 16:50:55'),
(5, NULL, 'username', 'password', NULL, NULL, 0, 0, 'user', '2022-06-25 16:51:11'),
(6, NULL, 'rafli', '123', NULL, NULL, 100, 2000, 'user', '2022-06-25 17:34:37'),
(7, 'urlpath', 'rafli', 'passworddiubah', '08151448223', 'iniemail', 0, 0, 'user', '2022-06-25 17:43:10'),
(8, NULL, 'rafli', '123', NULL, NULL, 0, 0, 'user', '2022-06-25 17:44:23');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
