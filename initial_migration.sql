CREATE SCHEMA `ecom_modo`;

CREATE TABLE `ecom_modo`.`buyers`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        datetime     DEFAULT current_timestamp,
    `updated_at`        datetime     DEFAULT current_timestamp,
    `deleted_at`        datetime     DEFAULT NULL,
    `email`             varchar(45)  DEFAULT NULL,
    `name`              varchar(45)  DEFAULT NULL,
    `password`          varchar(255) DEFAULT NULL,
    `alamat_pengiriman` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY                 `idx_buyers_deleted_at_email` (`deleted_at`, `email`)
) ENGINE=InnoDB;

CREATE TABLE `ecom_modo`.`sellers`
(
    `id`             int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime     DEFAULT current_timestamp,
    `updated_at`     datetime     DEFAULT current_timestamp,
    `deleted_at`     datetime     DEFAULT NULL,
    `email`          varchar(45)  DEFAULT NULL,
    `name`           varchar(45)  DEFAULT NULL,
    `password`       varchar(255) DEFAULT NULL,
    `alamat_pick_up` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY              `idx_sellers_deleted_at_email` (`deleted_at`, `email`)
) ENGINE=InnoDB;

CREATE TABLE `ecom_modo`.`products`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime       DEFAULT current_timestamp,
    `updated_at`  datetime       DEFAULT current_timestamp,
    `deleted_at`  datetime       DEFAULT NULL,
    `name`        varchar(100)   DEFAULT NULL,
    `description` varchar(255)   DEFAULT NULL,
    `price`       decimal(20, 3) DEFAULT NULL,
    `seller_id`   int unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY           `idx_products_deleted_at_seller_id` (`deleted_at`, `seller_id`)
) ENGINE=InnoDB;

CREATE TABLE `ecom_modo`.`orders`
(
    `id`                      int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`              datetime     DEFAULT current_timestamp,
    `updated_at`              datetime     DEFAULT current_timestamp,
    `deleted_at`              datetime     DEFAULT NULL,
    `buyer_id`                int unsigned DEFAULT NULL,
    `seller_id`               int unsigned DEFAULT NULL,
    `delivery_source_address` varchar(255) DEFAULT NULL,
    `delivery_dest_address`   varchar(255) DEFAULT NULL,
    `items`                   int unsigned DEFAULT NULL,
    `quantity`                int          DEFAULT NULL,
    `price`                   double       DEFAULT NULL,
    `total_price`             double       DEFAULT NULL,
    `status`                  varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY                       `idx_orders_deleted_at_seller_id` (`deleted_at`, `seller_id`),
    KEY                       `idx_orders_deleted_at_buyer_id` (`deleted_at`, `buyer_id`)
) ENGINE=InnoDB;

INSERT INTO `ecom_modo`.`buyers` (`email`, `name`, `password`, `alamat_pengiriman`)
VALUES ('test.buyer1@mail.com', 'test buyer 1', '$2a$10$HsQdPsxy8YSlQ.hOvurql.P5voRPOPWFvMaKAiLbZuHmEvHF6Zzi2',
        'jalan buntu');
INSERT INTO `ecom_modo`.`buyers` (`email`, `name`, `password`, `alamat_pengiriman`)
VALUES ('test.buyer2@mail.com', 'test buyer 2', '$2a$10$iosE4T0otbeidnApsAXcoOnU8DdWg1ws0YDS5SMV.OTdygQ6X7KC2',
        'jalan gang');
INSERT INTO `ecom_modo`.`buyers` (`email`, `name`, `password`, `alamat_pengiriman`)
VALUES ('test.buyer3@mail.com', 'test buyer 3', '$2a$10$969skMO.hiD6nCEqZORffu0vn4X0yUiWmAk2X9W2t0Ra/miiCh6DK',
        'jalan raya');

INSERT INTO `ecom_modo`.`sellers` (`email`, `name`, `password`, `alamat_pick_up`)
VALUES ('test.seller1@mail.com', 'test seller 1', '$2a$10$rD9Tc1OtROloZdaZywXfle..TNGot771jqxC154sQeS8.6vltbfFy',
        'rumah');
INSERT INTO `ecom_modo`.`sellers` (`email`, `name`, `password`, `alamat_pick_up`)
VALUES ('test.seller2@mail.com', 'test seller 2', '$2a$10$KcCh84Xbzbosup3kA7RBuuMUwSb5rTV8cpBiA4gIx.k80NQPgHOVu',
        'kantor');
INSERT INTO `ecom_modo`.`sellers` (`email`, `name`, `password`, `alamat_pick_up`)
VALUES ('test.seller3@mail.com', 'test seller 3', '$2a$10$PiqLPxIOis24KKl.a.bcc.G/4vuNXrv89Jli1M50iCkTSQStD6ug.',
        'apartemen');

INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('Indomie', 'mie instant', '5000.50', '1');
INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('Chiki', 'snack', '2000.00', '1');
INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('Laptop', 'elektronik', '1500000.50', '2');
INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('TV', 'elektronik', '1000000.00', '2');
INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('Jeans', 'fashion', '50000.50', '3');
INSERT INTO `ecom_modo`.`products` (`name`, `description`, `price`, `seller_id`)
VALUES ('T-shirt', 'fashion', '75000.50', '3');