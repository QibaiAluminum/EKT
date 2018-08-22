# blockchain_manager
## 文件内容及功能描述
* blockchain_manager.go
bug fixed: 转账到没有余额账户的bug

## 库
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
* blockchain.BlockChain
* blockchain.NewBlockChain
* blockchain.BackboneChainId
* blockchain.BackboneConsensus
* blockchain.BackboneChainFee
* blockchain.BackboneBlockInterval
* consensus.DbftConsensus
* consensus.NewDbftConsensus
* consensus.StableRun
* db.GetDBInst
* i_consensus.Consensus
* i_consensus.DBFT

## 对外暴露的方法
* blockchain_manager.MainBlockChainConsensus
* blockchain_manager.Init
