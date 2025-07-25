DROP TABLE IF EXISTS SkuCarouselPics;
CREATE TABLE SkuCarouselPics (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  sku_id TEXT,
  pic_str TEXT
);
INSERT INTO SkuCarouselPics (sku_id, pic_str) VALUES ("g7225946", "-01.jpeg");
INSERT INTO SkuCarouselPics (sku_id, pic_str) VALUES ("g7225946", "-02.jpeg");
INSERT INTO SkuCarouselPics (sku_id, pic_str) VALUES ("g7225947", "-01.jpeg");
INSERT INTO SkuCarouselPics (sku_id, pic_str) VALUES ("g7225947", "-02.jpeg");