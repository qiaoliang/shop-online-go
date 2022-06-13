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
INSERT INTO 
user_cart_items (Token,sku_Id,Pic,Status,Name,Sku_Strs,Price,Quantity,Selected,option_value_name)
VALUES 
('13900007997','g7225946','g7225946.jpeg',0,'持续交付1.0','sku1,sku3','66.0',110,'1','OptionValueName');