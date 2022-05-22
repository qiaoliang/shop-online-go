# API

URL 的前缀都是`http://{hostname:port}`

所有API均返回 `JSON` 格式的数据。

方法||URI|版本|说明|
|--|--|--|--|
|GET    |`/books`       |v1| fetch all books |
|GET    |`/books/${id}` |v1| fetch one book by id |
|POST   |`/books`       |v1| add a new book|
|PATCH  |`/books/${id}` |v1| update infomation for a book|
|DELETE |`/books/${id}` |v1| delete a book|