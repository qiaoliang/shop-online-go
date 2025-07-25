DROP TABLE IF EXISTS `addresses`;
CREATE TABLE `addresses` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` VARCHAR(100) NOT NULL,
    `link_man` VARCHAR(100) NOT NULL,
    `mobile` VARCHAR(100) NOT NULL,
    `province_str` VARCHAR(100) NOT NULL,
    `city_str` VARCHAR(100) NOT NULL,
    `area_str` VARCHAR(100) NOT NULL,
    `detail_address` VARCHAR(255) NOT NULL,
    `is_default` INTEGER NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`Id`)
);