# Go/Gin で Todo アプリを作る

## Library
* [Gin](https://github.com/gin-gonic/gin)

## Run application

```shell
go run src/main/main.go
```

or

```shell
go build src/main/main.go && ./main
```

## about Database

using `sqlite3`

### Library

* ORM: gorm
  - `github.com/jinzhu/gorm`
* Driver:
  - `github.com/mattn/go-sqlite3`

### Operations

* Database migration

```shell
go run src/main/db/migrate.go
```

## Link
* [Go / Gin で超簡単なWebアプリを作る](https://qiita.com/hyo_07/items/59c093dda143325b1859)