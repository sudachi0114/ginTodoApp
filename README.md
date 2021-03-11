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

### check by CLI client

* use `sqlite3`

```shell
# 起動
$ sqlite3
sqlite> 

# import database
sqlite3> open test.sqlite3

# database 一覧
sqlite> .databases

# table 一覧
sqlite> .tables

# todo table の schema を表示 (DESC)
sqlite> .schema todos

# todo table の中身を全部 SELECT
sqlite> SELECT * FROM todos;

# 終了
sqlite> .exit
# もしくは
sqlite> .quit
```

## Link
* [Go / Gin で超簡単なWebアプリを作る](https://qiita.com/hyo_07/items/59c093dda143325b1859)
* [SQLite3 Command Line Shell dot-commands マニュアル(完全版)](https://qiita.com/kanegoon/items/fc1e4bfea0984dbe4b90)
* [SQLiteをCLIから使う（sqlite3）](https://qiita.com/aki3061/items/f6450bdf3675418f0ef0)