DROP TABLE IF EXISTS banners;
CREATE TABLE banners (
  businessId INTEGER,
  dateAdd TEXT,
  id INTEGER PRIMARY KEY,
  linkUrl TEXT,
  paixu INTEGER,
  picUrl TEXT,
  remark TEXT,
  status INTEGER,
  statusStr TEXT,
  title TEXT,
  type TEXT,
  userId INTEGER
);
INSERT INTO banners (businessId, dateAdd, id, linkUrl, paixu, picUrl, remark, status, statusStr, title, type, userId) VALUES
(0, '2022-05-05 11:26:09', 222083, 'https://baidu.com', 0, 'http://localhost:9090/pic/banners/b0001.jpeg', '跳转gitee sagittatius', 0, 'any', 'any', 'any', 1605),
(1, '2022-05-05 11:26:09', 222084, 'https://baidu.com', 0, 'http://localhost:9090/pic/banners/b0002.jpeg', '跳转gitee sagittatius', 0, 'any', 'any', 'any', 1606);