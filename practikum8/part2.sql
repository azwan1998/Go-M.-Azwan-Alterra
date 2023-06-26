-- create new schema
CREATE SCHEMA `altra_online_shop` ; 


-- create table user
CREATE TABLE `altra_online_shop`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nama` VARCHAR(45) NOT NULL,
  `alamat` VARCHAR(255) NOT NULL,
  `tanggal_lahir` DATE NOT NULL,
  `status_user` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_ar` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));

-- create table product
CREATE TABLE `altra_online_shop`.`product` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `product_name` VARCHAR(255) NOT NULL,
  `product_description` VARCHAR(255) NOT NULL,
  `operator` VARCHAR(255) NOT NULL,
  `stock` INT NOT NULL,
  `price` INT NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));

-- create table transaksi
CREATE TABLE `altra_online_shop`.`transaksi` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_user` INT NOT NULL,
  `id_product` INT NOT NULL,
  `transaksi_description` VARCHAR(255) NOT NULL,
  `payment_method` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));

--create table shipping
CREATE TABLE `altra_online_shop`.`shipping` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_transaksi` INT NOT NULL,
  `created_at` TIMESTAMP(6) NULL,
  `updated_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`));

-- update table shipping add column description
ALTER TABLE `altra_online_shop`.`shipping` 
ADD COLUMN `description` VARCHAR(255) NULL AFTER `id_transaksi`;








