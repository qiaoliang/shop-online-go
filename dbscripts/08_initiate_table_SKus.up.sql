DROP TABLE IF EXISTS `skus`;
CREATE TABLE IF NOT EXISTS `skus` (
	`Sku_Id` VARCHAR(100) ,
	`Name` VARCHAR(100) NOT NULL,            
	`Category_Id` INTEGER NOT NULL,      
	`Recommend_Status` VARCHAR(100), 
	`Pic_Str` VARCHAR(100) NOT NULL,          
	`Unit` VARCHAR(100) NOT NULL,            
	`Stock` INTEGER ,           
	`Min_Price` VARCHAR(100) ,        
	`Original_Price` VARCHAR(100) NOT NULL,   
	`Logistics` VARCHAR(100),       
	`Content` VARCHAR(100) NOT NULL,         
	`Status` INTEGER,
	`Aftersale` INTEGER,
	PRIMARY KEY (Sku_Id)
);

INSERT INTO 
skus (`Sku_Id`,`Name`,`Category_Id`,`Recommend_Status`,`Pic_Str`,`Unit`,`Stock`,`Min_Price`,`Original_Price`,`Logistics`,`Content`,`Status`,`Aftersale`)
VALUES 
( "g7225946","持续交付1.0", 0, "1", "g7225946.jpeg", "册", 110, "66.0", "99.0", "1", "DevOps 的第一本书", 0, 1),
( "g7225947","持续交付2.0", 0, "1", "g7225947.jpeg", "册", 200, "99.0", "129.0", "1", "另一本DevOps的经典书。", 0, 1),
( "g7225948","DevOps实战指南", 0, "1", "g7225948.jpeg", "册", 10, "55.0", "85.0", "1", "DevOps 黄皮书。", 0, 1),
( "g7225949","谷歌软件工程", 0, "1", "g7225949.jpeg", "册", 20, "77.0", "107.0", "1", "解密硅谷头部互联网企业 如何打造软件工程文化。", 0, 1),
( "g1872110","驾驭大数据", 1, "1", "g1872110.jpeg", "册", 5, "110.0", "129.0", "1", "《驾驭大数据》是一本大数据的入门级好书。", 0, 1),
( "g1872111","数据分析变革", 1, "1", "g1872111.jpeg", "册", 15, "45.0", "99.0", "1", "《数据分析变革》是一本讲解如何通过数据分析指引业务增长与创新的书。", 0, 1),
( "g1872112","大数据测试技术与实践", 1, "1", "g1872112.jpeg", "册", 25,  "22.0", "85.0", "1", "《大数据测试技术与实践》是一本讲如何进行大数据系统测试与工具方法论的书。", 0, 1),
( "g1872113","图解Spark 大数据快速分析实战", 1, "1", "g1872113.jpeg", "册", 35,  "86.0", "107.0", "1", "《图解Spark 》是一本如何使用Spark进行大数据处理的书。", 0, 1);