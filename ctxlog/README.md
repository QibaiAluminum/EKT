# ctxlog
## 文件内容及功能描述
* ctxlog.go
重构contextLog的Finish逻辑，把拼装字符串暴露出去

## 库
* ctxlog.go
```
	"encoding/json"
	"fmt"
	"github.com/EducationEKT/EKT/log"
```

## 对下层的依赖
* log.Debug

## 对外暴露的方法
* ctxlog.Log
* ctxlog.Finish
* ctxlog.NewContextLog
* ctxlog.ContextLog
