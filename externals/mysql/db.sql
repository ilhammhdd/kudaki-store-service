CREATE DATABASE IF NOT EXISTS kudaki_store DEFAULT COLLATE = utf8_general_ci;

CREATE USER IF NOT EXISTS 'kudaki_user'@'localhost' IDENTIFIED BY 'kudakiuserrocks';

GRANT ALL PRIVILEGES ON kudaki_store.* TO 'kudaki_user'@'localhost'
WITH GRANT OPTION;

USE kudaki_store;

CREATE TABLE IF NOT EXISTS storefronts(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `user_uuid` VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS brands(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `name` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS items(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `brand_uuid` VARCHAR(64),
    `name` VARCHAR(255),
    `amount` INT(20),
    `unit` VARCHAR(255),
    `price` INT(20),
    `description` TEXT,
    `photo` VARCHAR(255),

    FOREIGN KEY(brand_uuid)
    REFERENCES brands(uuid)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS item_storefronts(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `storefront_uuid` VARCHAR(64),
    `item_uuid` VARCHAR(64),

    FOREIGN KEY(storefront_uuid)
    REFERENCES storefronts(uuid)
    ON DELETE CASCADE,
    FOREIGN KEY(item_uuid)
    REFERENCES items(uuid)
    ON DELETE CASCADE
);
