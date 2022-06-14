DROP TABLE IF EXISTS `user_cart_items`;
CREATE TABLE IF NOT EXISTS `user_cart_items` (
    `Token` VARCHAR(100) NOT NULL,
    `sku_id` VARCHAR(100) NOT NULL,
    `Pic` VARCHAR(100) NOT NULL,
    `Status`  INTEGER NOT NULL,
    `Name` VARCHAR(100) NOT NULL,
    `sku_strs` VARCHAR(100)NOT NULL,
    `Price` VARCHAR(100) NOT NULL,
    `Quantity` INTEGER NOT NULL,
    `Selected` VARCHAR(100) NOT NULL,
    `option_value_name` VARCHAR(100)NOT NULL
);