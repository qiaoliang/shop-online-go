DROP TABLE IF EXISTS skus;
CREATE TABLE skus (
  sku_id TEXT,
  name TEXT NOT NULL,
  category_id INTEGER NOT NULL,
  recommend_status TEXT,
  pic_str TEXT NOT NULL,
  unit TEXT NOT NULL,
  stock INTEGER,
  min_price TEXT,
  original_price TEXT NOT NULL,
  logistics TEXT,
  content TEXT NOT NULL,
  status INTEGER,
  aftersale INTEGER
);

INSERT INTO skus (sku_id, name, category_id, recommend_status, pic_str, unit, stock, min_price, original_price, logistics, content, status, aftersale)
VALUES ('g7225946', '持续交付1.0', 0, '1', 'g7225946.jpeg', '本', 10, '66.0', '99.0', '1', '这是第一本 DevOps 的书', 0, 1);
INSERT INTO skus (sku_id, name, category_id, recommend_status, pic_str, unit, stock, min_price, original_price, logistics, content, status, aftersale)
VALUES ('g7225947', '持续交付2.0', 0, '1', 'g7225947.jpeg', '本', 20, '88.0', '120.0', '1', '这是第二本 DevOps 的书', 0, 1);