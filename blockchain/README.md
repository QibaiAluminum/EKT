# blockchain
## 文件内容及功能描述
* block.go
blockchain增加处理不同类型userevent的方法
* block_manager.go
适配新的交易池和新的userevent模块
* block_record.go
refact code
* block_validator.go
refact code
* blockchain.go
blockchain增加处理不同类型userevent的方法
* body.go
修改从[]byte到blockBody的方法，增加代码可读性
* body_test.go
blockBody不再带有高度
* evil_block.go
refact code
* police.go
refact code
* vote.go
bug fixed: 转账到没有余额账户的bug

## 库
* block.go
```
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/EducationEKT/EKT/MPTPlus"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/round"
```
* block_manager.go
```
	"encoding/hex"
	"sync"
	"time"
```
* block_record.go
```
	"encoding/hex"
	"sync"
```
* block_validator.go
```
	"bytes"
	"fmt"
	"sync"
	"time"
```
* blockchain.go
```
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"errors"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/ctxlog"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/i_consensus"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/pool"
```
* body.go
```
	"encoding/json"
	"github.com/EducationEKT/EKT/core/userevent"
```
* body_test.go
```
	"encoding/json"
	"fmt"
	"testing"
```
* evil_block.go
```
	"encoding/json"
```
* police.go
```
	"bytes"
	"time"
```
* vote.go
```
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/p2p"
```
## 对下层的依赖
* crypto.Sha3_256
* crypto.RecoverPubKey
* crypto.Crypto
* ctxlog.ContextLog
* db.GetDBInst
* event.EventId
* event.Type
* i_consensus.DBFT
* i_consensus.ConsensusType
* log.Debug
* log.Info
* p2p.Peer
* pool.TxPool
* pool.NewMultiFatcher
* Pool.MultiFetcher
* round.Round
* types.Account
* types.NewAccount
* types.HexBytes
* types.FromPubKeyToAddress
* userevent.Transaction
* userevent.TxResult
* userevent.TokenIssue
* userevent.IUserEvent
* userevent.NewTransactionResult
* userevent.TYPE_USEREVENT_TRANSACTION
* MPTPlus.MTP
* MPTPlus.MTP_Tree
* MPTPlus.NewMTP
## 对外暴露的方法
* blockchain.BlockChain
* blockchain.NewBlockChain
* blockchain.BackboneChainId
* blockchain.BackboneConsensus
* blockchain.BackboneChainFee
* blockchain.BackboneBlockInterval     
* blockchain.BlockManager
* blockchain.Block
* blockchain.BlockVote
* blockchain.VoteResults
* blockchain.NewVoteResults
* blockchain.BLOCK_ERROR_START
* blockchain.BLOCK_ERROR_END
* blockchain.BLOCK_SAVED
* blockchain.BLOCK_VOTED
* blockchain.BLOCK_VALID
