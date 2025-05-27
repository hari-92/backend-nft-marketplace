-- Create "orders" table
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` int unsigned NOT NULL,
  `pair_id` int unsigned NOT NULL,
  `order_type` varchar(10) NOT NULL,
  `side` varchar(10) NOT NULL,
  `price` decimal(18,8) NOT NULL,
  `quantity` decimal(18,8) NOT NULL,
  `filled_quantity` decimal(18,8) NULL DEFAULT 0.00000000,
  `status` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_orders_deleted_at` (`deleted_at`),
  CONSTRAINT `chk_orders_order_type` CHECK (`order_type` in (_utf8mb4'LIMIT',_utf8mb4'MARKET')),
  CONSTRAINT `chk_orders_side` CHECK (`side` in (_utf8mb4'BUY',_utf8mb4'SELL')),
  CONSTRAINT `chk_orders_status` CHECK (`status` in (_utf8mb4'OPEN',_utf8mb4'PENDING',_utf8mb4'FILLED',_utf8mb4'CANCELLED'))
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "pairs" table
CREATE TABLE `pairs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `base_token_id` int unsigned NOT NULL,
  `quote_token_id` int unsigned NOT NULL,
  `symbol` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_pairs_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "pnls" table
CREATE TABLE `pnls` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` int unsigned NOT NULL,
  `pair_id` int unsigned NOT NULL,
  `realized_pnl` decimal(18,8) NOT NULL,
  `unrealized_pnl` decimal(18,8) NOT NULL,
  `calculated_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_pnls_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "tokens" table
CREATE TABLE `tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `address` varchar(42) NOT NULL,
  `symbol` varchar(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text NULL,
  `decimals` bigint NOT NULL,
  `total_supply` bigint NOT NULL,
  `chain_id` bigint NOT NULL,
  `is_active` bool NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `idx_tokens_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `username` varchar(50) NULL,
  `email` varchar(100) NOT NULL,
  `password_hash` varchar(255) NULL,
  `google_id` varchar(100) NULL,
  `phone_number` varchar(20) NULL,
  `lang` varchar(2) NULL DEFAULT "en",
  `is_set_lang` bool NULL DEFAULT 0,
  `last_login_at` datetime(3) NULL,
  `test` varchar(100) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`),
  UNIQUE INDEX `uni_users_email` (`email`),
  UNIQUE INDEX `uni_users_username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "wallets" table
CREATE TABLE `wallets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` int unsigned NOT NULL,
  `token_id` int unsigned NOT NULL,
  `balance` decimal(18,8) NOT NULL DEFAULT 0.00000000,
  `locked_balance` decimal(18,8) NOT NULL DEFAULT 0.00000000,
  `address` longtext NOT NULL,
  `chain_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_wallets_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
