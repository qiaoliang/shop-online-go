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

# API

URL 的前缀都是`http://${yourHostnameOrIP}:${port}`

方法||URI|版本|说明|
|--|--|--|--|
|GET    |`/books`                           |v1| fetch all books by json|
|GET    |`/books/${id}`                     |v1| fetch one book by id|
|POST   |`/books`                           |v1| add a new book|
|PATCH  |`/books/${id}`                     |v1| update infomation for a book|
|DELETE |`/books/${id}`                     |v1| delete a book|