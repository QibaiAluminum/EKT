# userevent
## 文件内容及功能描述
* iuserevent.go
* tokenissue.go
* transaction.go
* userevent_test.go

## 库
* iuserevent.go
```
	"bytes"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/crypto"
	"sort"
	"strings"
```
* tokenissue.go
```
	"encoding/hex"
	"encoding/json"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/crypto"
```
* transaction.go
```
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/db"
```
* userevent_test.go
```
	"fmt"
	"github.com/EducationEKT/EKT/crypto"
	"sort"
	"testing"
```

## 对下层的依赖
* crypto.Crypto
* crypto.RecoverPubKey
* crypto.Sha3_256
* db.GetDBInst
* types.FromPubKeyToAddress
* types.HexBytes
* types.Token

## 对外暴露的方法
* event.EventId
* event.Type
* userevent.GetTransaction
* userevent.TokenIssue
* userevent.IUserEvent
* userevent.NewTransaction
* userevent.Validate
* userevent.Transaction
* userevent.TYPE_USEREVENT_TRANSACTION
* userevent.TYPE_USEREVENT_PUBLIC_TOKEN

