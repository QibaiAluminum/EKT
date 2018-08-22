# conf
## 文件内容及功能描述
* conf.go
修改获取私钥的方法

## 库
```
	"encoding/json"
	"io/ioutil"

	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/p2p"
```

## 对下层的依赖
* types.Account
* types.HexBytes
* p2p.Peer

## 对外暴露的方法
* conf.InitConfig

