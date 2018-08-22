# db
## 文件内容及功能描述
* ComposedKVDatabase.go
* errors.go
* IKVDatabase.go
* init.go
* leveldb.go
* leveldb_test.go
* MemKVDatabase.go
* sqlite.go

## 库
* ComposedKVDatabase.go
* errors.go
```
"errors"
```
* IKVDatabase.go
* init.go
* leveldb.go
```
"github.com/syndtr/goleveldb/leveldb"
```
* leveldb_test.go
```
"testing"
```
* MemKVDatabase.go
```
	"encoding/hex"
	"sync"
```
* sqlite.go
```
	"database/sql"

	"github.com/mattn/go-sqlite3"
```

## 对下层的依赖
* leveldb.DB.Put
* levelDB.DB.Get
* levelDB.DB.Delete
* leveldb.OpenFile
* sql.Open

## 对外暴露的方法
* db.DB.Close
* db.GetDBInst
* db.InitEKTDB
* db.NewLevelDB

