# p2p
## 文件内容及功能描述
* peer_test.go
refact code
* types.go
修改peerId逻辑，不存在独立的peerId, 与用户的私钥和地址公用同一套逻辑

## 库
* peer_test.go
```
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/EducationEKT/EKT/crypto"
```
* types.go
```
	"encoding/json"
	"fmt"

	"bytes"
	"github.com/EducationEKT/EKT/util"
	"strings"
```

## 对下层的依赖
* crypto.Sha3_256
* util.HttpGet
* util.HttpPost

## 对外暴露的方法
* p2p.Peer


