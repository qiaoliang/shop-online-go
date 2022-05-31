# API

URL 的前缀都是`http://{hostname:port}`

所有API均返回 `JSON` 格式的数据。

方法|URI|版本|说明|
|--|--|--|--|
|GET    |`/books`       |v1| fetch all books |
|GET    |`/books/${id}` |v1| fetch one book by id |
|POST   |`/books`       |v1| add a new book|
|PATCH  |`/books/${id}` |v1| update infomation for a book|
|DELETE |`/books/${id}` |v1| delete a book|


## v1 for shop

### Verification 

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|GET     |`/v1/verification/pic/get` | ? |v1| fetch captCha |
|GET    |`/v1/verification/pic/check` |? |v1| Verify CapChar  |
|GET    |`/v1/verification/sms/get` |? |v1|  Get SMS Code  |




### Advertise

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|GET    |`/v1/banner/list` |? |v1| fetch advertise for banner  |

### User Management

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|POST    |`/v1//user/m/register` |?|v1| user register  |
|POST    |`/v1/user/m/login` |? |v1| user login  |
|GET    |`/v1/user/detail` |? |v1| Get User Detail  |
|GET    |`/v1/user/modify` |? |v1| Update User Info  |
|GET    |`/v1/user/amount` |? |v1| Get User Amount  |
|GET    |`/v1/user/logout` |? |v1| Get User Detail  |

### Order Management

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|GET    |`/v1/order/statistics` |? |v1| Discount Statistics  |
|GET    |`/v1/discounts/statistics` |? |v1| Get Order Statistics  |
|GET    |`/v1/discounts/coupons` |? |v1| Coupons  |

### Goods Management

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|GET    |`/v1/shop/goods/category/all` |? |v1| FetchCatalogues  |
|GET    |`/v1/shop/goods/detail` |? |v1| GetGoodsDetail  |
|GET    |`/v1/shop/goods/reputation` |? |v1| FetchItemReputation  |
|POST    |`/v1/goods/list",` |? |v1| FetchGoodsList  |

### Shopping Cart Management

方法|URI|参数|版本|说明|
|--|--|--|--|--|
|GET    |`/v1/shopping-cart/info` |? |v1| GetShopingCart  |
|POST    |`/v1/shopping-cart/add` |? |v1| PutIntoCart  |
|POST    |`/v1/shopping-cart/modifyNumber` |? |v1| UpdateShoppingCart  |

