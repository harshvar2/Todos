
CREATE TABLE `todos` (
  `id` bigint(20) UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'name associated with the todo',
  `description` VARCHAR(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'description of the todo',
  `status` enum('todo','in-progress','done') COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'status of the todo',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'Timestamp at which the row was created',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp at which the row was last updated',
    PRIMARY KEY (id)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `todos` COMMENT = 'Used for storing list of todos'; 
