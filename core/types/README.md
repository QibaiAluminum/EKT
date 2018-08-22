# types
## 文件内容及功能描述
* account.go
package refact
* address.go
package refact
* hexbytes.go
package refact
* hexbytes_test.go
package refact
* token.go
token增加序列化的方法
* types.go
package refact

## 库
* account.go
```
	"bytes"
	"encoding/json"
	"github.com/EducationEKT/EKT/crypto"
```
* address.go
```
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto/sha3"
```
* hexbytes.go
```
	"encoding/hex"
	"fmt"
```
* hexbytes_test.go
```
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
```
* token.go
```
	"encoding/json"
	"github.com/EducationEKT/EKT/crypto"
```
* types.go


## 对下层的依赖
* crypto.Sha3_256

## 对外暴露的方法
* types.Account
* types.HexBytes
* types.FromPubKeyToAddress
* types.Token.Name
* types.Token.Decimals 
* types.Token.Total
