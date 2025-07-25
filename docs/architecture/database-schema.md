# Database Schema

```sql
-- Table: addresses
CREATE TABLE IF NOT EXISTS `addresses` (
    `id` VARCHAR(36) NOT NULL PRIMARY KEY COMMENT '地址唯一标识',
    `user_id` VARCHAR(36) NOT NULL COMMENT '关联的用户 ID',
    `link_man` VARCHAR(255) NOT NULL COMMENT '联系人姓名',
    `mobile` VARCHAR(20) NOT NULL COMMENT '联系手机号',
    `province_str` VARCHAR(255) NOT NULL COMMENT '省份',
    `city_str` VARCHAR(255) NOT NULL COMMENT '城市',
    `area_str` VARCHAR(255) NOT NULL COMMENT '区域',
    `detail_address` VARCHAR(500) NOT NULL COMMENT '详细地址',
    `is_default` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否为默认地址',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户收货地址表';

-- Foreign Key (假设 users 表已存在)
-- ALTER TABLE `addresses`
-- ADD CONSTRAINT `fk_addresses_user_id`
-- FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
-- ON DELETE CASCADE ON UPDATE CASCADE;

-- 索引用于快速查找用户的默认地址
CREATE INDEX `idx_user_id_is_default` ON `addresses` (`user_id`, `is_default`);
```
