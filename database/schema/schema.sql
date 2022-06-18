DROP TABLE IF EXISTS `brands`;
CREATE TABLE `brands` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) unique,
    `description` varchar(255),
    `logo` varchar(255),
    `level` enum('small', 'medium', 'large'),
    `is_active` boolean default true,
    `created_at` timestamp default now()
);

CREATE INDEX brands_index_level ON `brands` (level);
CREATE INDEX brands_index_name_is_active ON `brands` (level, is_active);

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `brand_id` bigint,
    `title` varchar(255) unique,
    `description` varchar(255),
    `price` bigint,
    `price_reduction` bigint,
    `stock` int default 0,
    `is_active` boolean default true,
    `created_at` timestamp default now()
);

ALTER TABLE `products` ADD FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`);
CREATE INDEX products_index_title_is_active ON `products` (title, is_active);

DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `state` varchar(255),
    `created_at` timestamp default now(),
    `updated_at` timestamp
);
CREATE INDEX brands_index_state ON `transactions` (state);

DROP TABLE IF EXISTS `transaction_details`;
CREATE TABLE `transaction_details` (
       `id` bigint PRIMARY KEY AUTO_INCREMENT,
       `transaction_id` bigint,
       `product_id` bigint,
       `price` bigint,
       `price_reduction` bigint,
       `final_price` bigint,
       `quantity` int default 1,
       `created_at` timestamp default now()
);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);
ALTER TABLE `transaction_details` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
