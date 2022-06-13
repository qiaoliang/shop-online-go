DROP TABLE IF EXISTS `UserCartItems`;
CREATE TABLE IF NOT EXISTS `UserCartItems` (
    `ID` INTEGER NOT NULL AUTO_INCREMENT,
    `Token` VARCHAR(100) NOT NULL,
    `sku_id` VARCHAR(100) NOT NULL,
    `Pic` VARCHAR(100) NOT NULL,
    `Status`  INTEGER NOT NULL,
    `Name` VARCHAR(100) NOT NULL,
    `SkuStrs` VARCHAR(100)NOT NULL,
    `Price` VARCHAR(100) NOT NULL,
    `Quantity` INTEGER NOT NULL,
    `Selected` VARCHAR(100) NOT NULL,
    `OptionValueName` VARCHAR(100)NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (Sku_Id) REFERENCES skus(Sku_Id)
);
INSERT INTO 
UserCartItems (Id,Token,sku_id,Pic,Status,Name,SkuStrs,Price,Quantity,Selected,OptionValueName)
VALUES 
(0, 'token-13900007997','g7225946','g7225946.jpeg',0,'持续交付1.0','sku1,sku3','66.0',110,'1','OptionValueName');