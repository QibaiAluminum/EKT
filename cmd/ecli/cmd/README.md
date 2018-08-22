# ecli_cmd
## 文件内容及功能描述
* account.go
命令行工具适配新的userevent
* nodes.go
删除node相关的子命令，新增发行token的命令
* transaction.go
删除node相关的子命令，新增发行token的命令

## 库
* account.go
```
	"encoding/hex"
	"fmt"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/spf13/cobra"
```
* nodes.go
```
	"encoding/hex"
	"fmt"

	"github.com/EducationEKT/EKT/crypto"
	"github.com/spf13/cobra"
```
* transaction.go
```
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/EducationEKT/EKT/cmd/ecli/param"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/util"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
```

## 对下层的依赖
* cobra.Command
* crypto.GenerateKeyPair
* crypto.Sha3_256
* crypto.PubKey
* param.GetPeers
* userevent.NewTransaction
* util.HttpPost
* util.HttpGet
* x_resp.XRespBody

## 对外暴露的方法
* cmd.TransactionCmd
* cmd.AccountCmd
* cmd.NodeCommand

