DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id TEXT PRIMARY KEY,
  pwd TEXT NOT NULL,
  mobile TEXT NOT NULL,
  nickname TEXT,
  avatar_url TEXT,
  province TEXT,
  city TEXT,
  auto_login INTEGER,
  user_info TEXT,
  User_Level_Id INTEGER
);
INSERT INTO users (id, mobile, pwd, nickname, avatar_url, province, city, auto_login, user_info, User_Level_Id)
VALUES ('admin', '13900007997', '1234', 'admin', 'a.jpeg', '未知', '未知', 1, 'FakeUserInfo', 1);