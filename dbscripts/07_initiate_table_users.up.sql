DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
	Id   VARCHAR(100)  PRIMARY KEY,
	Password   VARCHAR(100) NOT NULL,
	Mobile   VARCHAR(100),
	Nickname   VARCHAR(100),
	AvatarUrl   VARCHAR(255),
	Province   VARCHAR(100),
	City   VARCHAR(100),
	AutoLogin  INTEGER,
	UserInfo   VARCHAR(100),
	UserLevelId   INTEGER
);

INSERT INTO 
users (Id,Password,Mobile,Nickname,AvatarUrl,Province,City,AutoLogin,UserInfo,UserLevelId)
VALUES 
('13900007997', '1234','13900007997','天之骄子', 'a.jpeg', '北京',  '北京',  1, '这是UserInfo',1);