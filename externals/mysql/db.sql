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

CREATE TABLE IF NOT EXISTS item_storefronts(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `storefront_uuid` VARCHAR(64),
    `item_uuid` VARCHAR(64),

    FOREIGN KEY(storefront_uuid)
    REFERENCES storefronts(uuid)
);

-- mamak japar : 085880692628