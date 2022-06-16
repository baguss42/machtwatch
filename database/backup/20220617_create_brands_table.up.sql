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