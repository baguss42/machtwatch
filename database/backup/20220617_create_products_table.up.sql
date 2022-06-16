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