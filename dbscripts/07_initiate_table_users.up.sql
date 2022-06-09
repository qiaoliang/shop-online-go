DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
	Id   VARCHAR(100),
	Password   VARCHAR(100) NOT NULL,
	Mobile   VARCHAR(100),
	Nickname   VARCHAR(100),
	avatar_url   VARCHAR(255),
	Province   VARCHAR(100),
	City   VARCHAR(100),
	auto_login  INTEGER,
	user_info   VARCHAR(100),
	User_Level_Id   INTEGER,
	PRIMARY KEY (Id)
);

INSERT INTO 
users (Id,Password,Mobile,Nickname,avatar_url,Province,City,auto_login,user_info,User_Level_Id)
VALUES 
('13900007997', '1234','13900007997','天之骄子', 'a.jpeg', '北京',  '北京',  1, '这是UserInfo',1);