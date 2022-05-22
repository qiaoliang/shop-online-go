# Book Store

Simple CRUD with [Gin](https://github.com/gin-gonic/gin) Framework

# Stack

- [GO](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/index.html)
- [Gin](https://github.com/gin-gonic/gin)
- [Migrate](https://github.com/golang-migrate)

# How To Run

1. copy or change `config.example.yaml` to `config.yaml`
2. config db info in `config.yml`
3. create mysql database instance with name `bookstore` by `CREATE DATABASE bookstore;` before running. 
4. run project `go run main.go` which will initial data for you by scripts under the dir `dbscripts`.

# How to verify

visit `http://${yourIP}:${PORT_in_config.yaml}/books`, you will see:

```
{"data":[{"id":1,"title":"little prince","author":"Antoine"},{"id":2,"title":"Les Trois Mousquetaires","author":"Alexandre Dumas fils"},{"id":3,"title":"Continuous Delivery","author":"Jez"}]}
```

# directory structure
* `app`: main code 
  * `configs`: Parser for config.yaml
  * `controllers`:  Implementation of the restful APIs 
  * `models`:  Data model
  * `routers`:  setup routers for the program
* `dbscripts`:  DB scripts for migration. It is invoked by main.go
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