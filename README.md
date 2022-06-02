# Book Store

Simple CRUD with [Gin](https://github.com/gin-gonic/gin) Framework

# Stack

- [GO](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/index.html)
- [Gin](https://github.com/gin-gonic/gin)
- [Migrate](https://github.com/golang-migrate)

# How To Run

1. config db info in `config.yml`
2. create mysql database instance with name `bookstore` by `CREATE DATABASE bookstore;` before running. 
4. run project `go run main.go` which will initial data for you by scripts under the dir `dbscripts`.

# How to verify

visit `http://${yourIP}:${PORT_in_config.yaml}/v1/shop/goods/category/all`, you will see:

```
{"code":0,"data":[{"id":0,"name":"DevOps"},{"id":1,"name":"大数据"}],"msg":"OK"}
```

# directory structure
* `app`: main code 
  * `configs`: Parser for config.yaml
  * `ad`:  Implementation of domain `Advertise management` 
  * `configs`:  Implementation of Domain `system configuration` 
  * `goods`:  Implementation of domain `goods and category` 
  * `order`:  Implementation of domain `Order management` 
  * `routers`:  setup routers for the program
  * `security`:  Implementation of domain `verification management` 
  * `shoppingcart`:  Implementation of domain `shoppingcate management`
  * `user`:  Implementation of domain `userinfo management`
  * `utils`:  Implementation of internal utils
* `dbscripts`:  DB scripts for migration. It is invoked by main.go
* `resources`:  font for captCha
* `static`:  static pictures for goods, avatars and ad banners
* `config.yaml`: Config for the program, mainly for database
* `main.go`: Entry of the program  

# 应用初始化

* 数据不支持持久化，即：每次重新启动，所有数据重置。
* 初始用户名为：13900007997，密码为：1234
* SKU只有两类，分别是 `DevOps` 和 `大数据`
* `DevOps`下有四个SKU， SKU的Id顺序为：g7225946～9，对应的图片保存在 `static` 目录下
* `大数据`下也有四个SKU， SKU的Id顺序为：g1872110～3，对应的图片保存在 `static` 目录下

# 业务规则

1. 「购物车」的红点一直显示，直到打开「购物车」的页面（与京东商城的逻辑一致）
2. 「购物车」的红点数等于该登录用户购物车中的 SKU 总数（与京东商城的逻辑一致）
3. 每次用户重新登录，只要「购物车」中有 SKU，就显示红点。