# dispatcher
## 文件内容及功能描述
* dispatcher.go
在交易池中使用指针而不是值传递

## 库
* dispatcher.go
```
	"encoding/hex"
	"errors"

	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/core/types"
	"github.com/EducationEKT/EKT/core/userevent"
```

## 对下层的依赖
* blockchain_manager.GetMainChain().GetLastBlock()
* types.Token.Name
* types.Token.Decimals 
* types.Token.Total
* userevent.Transaction
* userevent.Validate

## 对外暴露的方法

