CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)  DEFAULT NULL,
    `updated_at` datetime(3)  DEFAULT NULL,
    `deleted_at` datetime(3)  DEFAULT NULL,
    `name`       varchar(255)    NOT NULL,
    `email`      varchar(255) DEFAULT NULL,
    `age`        int          DEFAULT NULL,
    `is_active`  tinyint(1)   DEFAULT '1',
    `points`     int          DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_users_email` (`email`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;