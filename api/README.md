# API
## 文件内容及功能描述
* block.go
删除无用日志
* db.go
删除无用api
* peer.go
将共识算法修改为dbft
* token.go
增加发行token的api
* transaction.go
将共识算法修改为dbft
* user.go
重构交易池和用户事件逻辑，优化TPS
* vote.go
使用新的uservent和consensus

## 库
* block.go
```
	"encoding/json"
	"fmt"
	"strings"
	"github.com/EducationEKT/EKT/blockchain"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/ctxlog"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/util"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* db.go
```
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/log"

	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* peer.go
```
	"github.com/EducationEKT/EKT/blockchain_manager"

	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* token.go
```
	"encoding/hex"
	"encoding/json"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* transaction.go
```
	"encoding/json"
	"fmt"
	"strings"

	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/dispatcher"
	"github.com/EducationEKT/EKT/param"
	"github.com/EducationEKT/EKT/util"

	"encoding/hex"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* user.go
```
	"encoding/hex"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/ctxlog"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```
* vote.go
```
	"encoding/json"

	"github.com/EducationEKT/EKT/blockchain"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
```

## 对下层的依赖
* blockchain_manager.MainBlockChainConsensus
* blockchain_manager.GetMainChain()
* ctxlog.Log
* ctxlog.Finish
* ctxlog.NewContextLog
* db.GetDBInst
* param.MainChainDelegateNode
* userevent.GetTransaction
* userevent.TokenIssue
* util.HttpPost
* x_err.XErr
* x_err.New
* x_err.NewXErr
* x_router.Post
* x_router.Get
* x_router.All
* x_resp.Return
* x_resp.Fail
* x_resp.XRespContainer
* x_resp.Success

## 对外暴露的方法
