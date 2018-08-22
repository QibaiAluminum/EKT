# pool
## 文件内容及功能描述
* pool.go
使用channel重构transaction pool，增加存储所有userevent的map
* pool_test.go
重构交易池和用户事件逻辑，优化TPS
* TxPool.go
bug fixed

## 库
* pool.go
```
	"encoding/hex"
	"github.com/EducationEKT/EKT/core/userevent"
	"sort"
	"sync"
```
* pool_test.go

* TxPool.go
```
	"encoding/hex"
	"github.com/EducationEKT/EKT/core/userevent"
```

## 对下层的依赖
* userevent.SortedUserEvent
需要创建新的信道调用
* userevent.IUserEvent
需要创建新的信道调用
* event.GetFrom
* event.GetNonce

## 对外暴露的方法
* pool.TxPool
* pool.NewTxPool
* pool.NewMultiFatcher
* pool.NewEventGetter
* Pool.MultiFetcher
