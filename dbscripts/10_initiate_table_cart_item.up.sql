DROP TABLE IF EXISTS user_cart_items;
CREATE TABLE user_cart_items (
  id INTEGER PRIMARY KEY,
  token TEXT NOT NULL,
  sku_id TEXT NOT NULL,
  pic TEXT NOT NULL,
  status INTEGER NOT NULL,
  name TEXT NOT NULL,
  sku_strs TEXT NOT NULL,
  price TEXT NOT NULL,
  quantity INTEGER NOT NULL,
  selected TEXT NOT NULL,
  option_value_name TEXT NOT NULL
);