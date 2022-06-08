DROP TABLE IF EXISTS `goods`;
CREATE TABLE IF NOT EXISTS `goods` (

	Gid  VARCHAR(100)  PRIMARY KEY, 
	GName VARCHAR(100) NOT NULL, 
	Pics  VARCHAR(256),
	CategoryId  INTEGER ,
    RecommendStatus VARCHAR(10),
	Stock   INTEGER,  
	Unit    VARCHAR(100), 
	Logistics INTEGER,  
	Content   VARCHAR(256), 
	CurState     INTEGER,   
	StatusStr    VARCHAR(10), 
	PicUrl       VARCHAR(256), 
	MinPrice      VARCHAR(256),
	OriginalPrice VARCHAR(256),
	AfterSale     INTEGER
);