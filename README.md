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
  * `ad`:  Implementation of domain `Advertise management` 
  * `addresses`:  Implementation of domain `Shipping Addresses management`. Not Implemented yet.
  * `configs`:  Implementation of Domain `system configuration` 
  * `goods`:  Implementation of domain `goods and category` 
  * `order`:  Implementation of domain `Order management`. Not Implemented yet. 
  * `routers`:  setup routers for the program
  * `security`:  Implementation of domain `verification management`. Not Implemented yet. 
  * `shoppingcart`:  Implementation of domain `shoppingcate management`
  * `user`:  Implementation of domain `userinfo management`
  * `utils`:  Implementation of internal utils
* `dbscripts`:  DB scripts for migration. It is invoked by main.go
* `resources`:  font for captCha
* `static`:  static pictures for goods, avatars and ad banners. This should be in CDN, but not implemented yet.
* `config.yaml`: Config for the program, mainly for database
* `main.go`: Entry of the program  

# 应用初始化

* Data for Goods/Category/ShoppingCart/User can be persistanted. 
* Memory version for data is still avalible by setting `persistance` as `false` in the fiile `config.yaml`, which will reset all data after every start.
* Username is `13900007997`，Passwrod is `1234`
* there are two categories of SKUs which are `DevOps` and `big data`.
  * `DevOps` has 4 SKUs and their IDs are between `g7225946` and `g7225949`
  * `big data` has 4 SKUs and their IDs are between `g1872110` and `g1872113`
* All images store under the folder `static`, named by SKU Id

# some rules

1. The reddot is always `ON` in「Shopping cart」until the page 「Shopping cart 」was opened. The rule follows JD.
2. the number of reddot is the total numbers of SKU in the user's 「shopping cart」.This rule follows JD too.
3. Given a user's shopping cart has SKUs stored, every time when he/she login, the reddot will show up.

# 自动化测试

* In order to seperate Testing Env from Prod Env, it will load testing config from the file `config-test.yaml` for testing，rather than `config.yaml`.
* You have to create database instance `bookstore` on MySQL before running test.
* run cmd `go test ./...`