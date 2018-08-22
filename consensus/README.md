# consensus
## 文件内容及功能描述
* dbft.go
修改从peer获取userevent的方法

## 库
* dbft.go
```
	"encoding/hex"
	"encoding/json"
	"fmt"
	"xserver/x_http/x_resp"
	"sync"
	"time"
	"bytes"
	"errors"
	"github.com/EducationEKT/EKT/MPTPlus"
	"github.com/EducationEKT/EKT/blockchain"
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/ctxlog"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/p2p"
	"github.com/EducationEKT/EKT/param"
	"github.com/EducationEKT/EKT/pool"
	"github.com/EducationEKT/EKT/round"
	"github.com/EducationEKT/EKT/util"
	"os"
	"runtime"
	"strings"
```

## 对下层的依赖
* blockchain.BLOCK_ERROR_START
* blockchain.BLOCK_ERROR_END
* blockchain.BLOCK_SAVED
* blockchain.BLOCK_VOTED
* blockchain.BLOCK_VALID
* blockchain.BlockChain
* blockchain.BlockManager
* blockchain.Block
* blockchain.BlockVote
* blockchain.VoteResults
* blockchain.NewVoteResults
* blockchain.NewBlockBody
* blockchain.FromBytes2BLockBody

* ctxlog.ContextLog
* ctxlog.Log
* ctxlog.Finish
* ctxlog.NewContextLog
* crypto.Sha3_256
* crypto.RecoverPubKey
* db.GetDBInst
* log.Crit
* log.Debug
* log.Info
* MPTPlus.NewMTP
* param.MainChainDelegateNode
* p2p.Peer
* pool.NewEventGetter
* types.FromPubKeyToAddress
* userevent.IUserEvent
* userevent.TokenIssue
* userevent.Validate
* userevent.Transaction
* userevent.TYPE_USEREVENT_TRANSACTION
* userevent.TYPE_USEREVENT_PUBLIC_TOKEN
* util.HttpPost
* util.HttpGet
* x_resp.XRespBody

## 对外暴露的方法
* consensus.DbftConsensus
* consensus.NewDbftConsensus
