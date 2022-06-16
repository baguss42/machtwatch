CREATE TABLE `brands` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `logo` varchar(255),
  `level` tinyint,
  `is_active` boolean,
  `created_at` timestamp
);

CREATE UNIQUE INDEX "brands_unique_index_name" ON `brands` ("name");
CREATE INDEX "brands_index_level" ON `brands` ("level");
CREATE INDEX "brands_index_name_is_active" ON `brands` ("level", "is_active");

CREATE TABLE `products` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `brand_id` bigint,
  `title` varchar(255),
  `description` varchar(255),
  `price` bigint,
  `price_reduction` bigint,
  `stock` int,
  `is_active` boolean,
  `created_at` timestamp
);

ALTER TABLE `products` ADD FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`);
CREATE UNIQUE INDEX "products_unique_index_title" ON `products` ("title");
CREATE INDEX "products_index_title_is_active" ON `products` ("title", "is_active");

CREATE TABLE `transactions` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `state` varchar(255),
  `total_price` bigint,
  `created_at` timestamp,
  `updated_at` timestamp
);
CREATE INDEX "brands_index_state" ON `transactions` ("state");

CREATE TABLE `transaction_details` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `transaction_id` bigint,
  `product_id` int,
  `price` bigint,
  `reduction_price` bigint,
  `final_price` bigint,
  `created_at` timestamp
);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);
ALTER TABLE `transaction_details` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
