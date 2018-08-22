# param
## 文件内容及功能描述
* bootnode.go
修改peerId逻辑，不存在独立的peerId, 与用户的私钥和地址公用同一套逻辑
* localnet.go
修改peerId逻辑，不存在独立的peerId, 与用户的私钥和地址公用同一套逻辑
* mainnet.go
refact code
* testnet.go
修改peerId逻辑，不存在独立的peerId, 与用户的私钥和地址公用同一套逻辑

## 库
* bootnode.go
```
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/p2p"
```
* localnet.go
```
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/EducationEKT/EKT/p2p"
```
* mainnet.go
```
	"github.com/EducationEKT/EKT/p2p"
```
* testnet.go
```
	"github.com/EducationEKT/EKT/p2p"
```

## 对下层的依赖
* conf.EKTConfig.Env
* p2p.Peer

## 对外暴露的方法
* param.MainChainDelegateNode
* param.InitBootNodes
* param.GetPeers
