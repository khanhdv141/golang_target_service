DROP TABLE IF EXISTS `documents`;
CREATE TABLE `documents` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `source_file_id` longtext,
  `preview_file_id` longtext,
  `editable_file_id` longtext,
  `metadata` json DEFAULT NULL,
  `created_by` longtext,
  `code` longtext,
  `type` longtext,
  `issuance_date` datetime(3) DEFAULT NULL,
  `publication_date` datetime(3) DEFAULT NULL,
  `expiration_date` datetime(3) DEFAULT NULL,
  `effective_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_documents_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` varchar(191) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `original_name` longtext,
  `mimetype` longtext,
  `size` bigint DEFAULT NULL,
  `path` varchar(300) DEFAULT NULL,
  `extension` longtext,
  `type` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_files_name` (`name`),
  UNIQUE KEY `idx_files_path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext,
  `password` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;