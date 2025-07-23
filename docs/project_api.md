# 接口文档

本项目API接口按照业务功能模块进行划分，方便开发人员快速查找和理解。

## 1. 安全模块 (security)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /v1/verification/pic/get | 获取图形验证码 | 无 | 返回图形验证码图片URL和验证码ID |
| GET /v1/verification/pic/check | 校验图形验证码 | `id`: string (验证码ID), `verifyCode`: string (用户输入的验证码) | `message`: string (校验结果信息) |
| GET /v1/verification/sms/get | 获取短信验证码 | `mobile`: string (手机号) | `message`: string (发送结果信息) |

## 2. 广告管理 (banner)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /v1/banner/list | 获取Banner列表 | 无 | `data`: array (Banner信息列表) |

## 3. 用户管理 (user)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| POST /v1/user/m/register | 用户注册 | `mobile`: string, `password`: string, `verifyCode`: string | `message`: string (注册结果信息) |
| POST /v1/user/m/login | 用户登录 | `mobile`: string, `password`: string | `token`: string (认证Token), `message`: string (登录结果信息) |
| GET /v1/user/detail | 获取用户详情 | 无 (通过Token认证) | `data`: object (用户详细信息) |
| GET /v1/user/modify | 修改用户信息 | `nickname`: string (可选), `avatarUrl`: string (可选), `province`: string (可选), `city`: string (可选) | `message`: string (修改结果信息) |
| GET /v1/user/amount | 获取用户资产 | 无 (通过Token认证) | `amount`: number (用户资产金额) |
| GET /v1/user/logout | 用户登出 | 无 (通过Token认证) | `message`: string (登出结果信息) |

## 4. 用户收货地址管理 (user/shipping-address)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| POST /v1/user/shipping-address/list | 获取用户收货地址列表 | 无 (通过Token认证) | `data`: array (收货地址列表) |
| GET /v1/user/shipping-address/default | 获取默认收货地址 | 无 (通过Token认证) | `data`: object (默认收货地址信息) |
| POST /v1/user/shipping-address/add | 添加收货地址 | `linkMan`: string, `mobile`: string, `provinceStr`: string, `cityStr`: string, `areaStr`: string, `detailAddress`: string, `isDefault`: bool (可选) | `message`: string (添加结果信息) |

## 5. 订单管理 (order)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /v1/order/statistics | 获取订单统计信息 | 无 (通过Token认证) | `data`: object (订单统计数据) |
| GET /v1/discounts/statistics | 获取折扣统计信息 | 无 (通过Token认证) | `data`: object (折扣统计数据) |
| GET /v1/discounts/coupons | 获取用户优惠券 | 无 (通过Token认证) | `data`: array (优惠券列表) |

## 6. 商品管理 (goods)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /v1/shop/goods/category/all | 获取所有商品分类 | 无 | `data`: array (商品分类列表) |
| GET /v1/shop/goods/detail | 获取商品详情 | `id`: string (商品ID) | `data`: object (商品详细信息) |
| POST /v1/shop/goods/reputation | 获取商品评价 | `goodsId`: string (商品ID) | `data`: array (商品评价列表) |
| POST /v1/goods/list | 获取商品列表 | `categoryId`: uint (可选), `name`: string (可选), `page`: int, `size`: int | `data`: array (商品列表), `total`: int (总数) |

## 7. 购物车管理 (shoppingcart)

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /v1/shopping-cart/info | 获取购物车信息 | 无 (通过Token认证) | `data`: object (购物车详情，包含商品列表) |
| POST /v1/shopping-cart/add | 添加商品到购物车 | `skuId`: string, `quantity`: uint | `message`: string (添加结果信息) |
| POST /v1/shopping-cart/modifyNumber | 修改购物车商品数量 | `skuId`: string, `quantity`: uint | `message`: string (修改结果信息) |

## 8. 书籍管理 (books) - 独立API组

| API名称 | 功能说明 | 参数数据说明 | 返回值说明|
|-----|---|----|----|
| GET /books | 获取所有书籍 | 无 | `data`: array (书籍列表) |
| POST /books | 创建新书籍 | `title`: string, `author`: string | `data`: object (新创建书籍的信息) |
| GET /books/:id | 获取单本书籍详情 | `id`: string (书籍ID) | `data`: object (书籍详细信息) |
| PATCH /books/:id | 更新书籍信息 | `id`: string (书籍ID), `title`: string (可选), `author`: string (可选) | `data`: object (更新后的书籍信息) |
| DELETE /books/:id | 删除书籍 | `id`: string (书籍ID) | `message`: string (删除结果信息) |
