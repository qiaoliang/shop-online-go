DROP TABLE IF EXISTS `skus`;
CREATE TABLE IF NOT EXISTS `skus` (
	Id VARCHAR(100)   PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,            
	Category_Id INTEGER NOT NULL,      
	Recommend_Status VARCHAR(100), 
	Pic_Str VARCHAR(100) NOT NULL,          
	Unit VARCHAR(100) NOT NULL,            
	Stock INTEGER ,           
	Min_Price VARCHAR(100) ,        
	Original_Price VARCHAR(100) NOT NULL,   
	Logistics VARCHAR(100),       
	Content VARCHAR(100) NOT NULL,         
	Status INTEGER,
	After_Sale VARCHAR(100) NOT NULL
);

INSERT INTO 
skus (Id,Name,Category_Id,Recommend_Status,Pic_Str,Unit,Stock,Min_Price,Original_Price,Logistics,Content,Status,After_Sale)
VALUES 
( "g7225946","持续交付1.0", 0, "1", "g7225946.jpeg", "册", 20, "66.0", "99.0", "1", "DevOps 的第一本书", "1", "1");
