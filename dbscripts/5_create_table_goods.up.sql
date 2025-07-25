DROP TABLE IF EXISTS goods;
CREATE TABLE goods (
  Gid  TEXT PRIMARY KEY,
  GName TEXT NOT NULL,
  Pics  TEXT,
  CategoryId  INTEGER,
  RecommendStatus TEXT,
  Stock   INTEGER,
  Unit    TEXT,
  Logistics INTEGER,
  Content   TEXT,
  CurState     INTEGER,
  StatusStr    TEXT,
  PicUrl       TEXT,
  MinPrice      TEXT,
  OriginalPrice TEXT,
  AfterSale     INTEGER
);