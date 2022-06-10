DROP TABLE IF EXISTS `SkuCarouselPics`;
CREATE TABLE IF NOT EXISTS `SkuCarouselPics` (
    `ID` int NOT NULL AUTO_INCREMENT,
	`Sku_Id` VARCHAR(100) ,
    `Pic_Str` VARCHAR(100) ,
    PRIMARY KEY (ID),
    FOREIGN KEY (Sku_Id) REFERENCES skus(Sku_Id)
);
INSERT INTO 
SkuCarouselPics (`Sku_Id`,`Pic_Str`)
VALUES 
( "g7225946", "-01.jpeg,-02.jpeg"),
( "g7225947", "-01.jpeg,-02.jpeg"),
( "g7225948", "-01.jpeg,-02.jpeg"),
( "g7225949", "-01.jpeg,-02.jpeg"),
( "g1872110", "-01.jpeg,-02.jpeg"),
( "g1872111", "-01.jpeg,-02.jpeg"),
( "g1872112", "-01.jpeg,-02.jpeg"),
( "g1872113", "-01.jpeg,-02.jpeg");