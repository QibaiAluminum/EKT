# MPTPlus
## 文件内容及功能描述
* trie.go
默克尔字典树使用IKVDatabase接口代替直接使用结构体
* trie_test.go
refact code
* types.go
adaptor package refact

## 库
* trie.go
```
	"bytes"
	"encoding/json"
	"errors"
	"github.com/EducationEKT/EKT/crypto"
```
* trie_test.go
```
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"
	"xserver/x_utils/x_random"
	"github.com/EducationEKT/EKT/db"
```
* types.go
```
	"sort"
	"sync"
	"bytes"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/db"
```

## 对下层的依赖
* crypto.Sha3_256
* db.NewLevelDB
* db.DB.Close
* db.IKVDatabase
* types.HexBytes
* x_random.RandomNumber

## 对外暴露的方法
* MPTPlus.MTP
* MPTPlus.MTP_Tree
