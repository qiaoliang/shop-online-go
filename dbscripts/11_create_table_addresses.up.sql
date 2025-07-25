DROP TABLE IF EXISTS addresses;
CREATE TABLE addresses (
  id INTEGER PRIMARY KEY,
  user_id TEXT NOT NULL,
  link_man TEXT NOT NULL,
  mobile TEXT NOT NULL,
  province_str TEXT NOT NULL,
  city_str TEXT NOT NULL,
  area_str TEXT NOT NULL,
  detail_address TEXT NOT NULL,
  is_default INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(Id)
);
-- 省略索引/约束部分，如有需要可补充