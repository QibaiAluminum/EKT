# xlog
## 文件内容及功能描述
* dailywriter.go
refact code
* xlog.go
xlog auto create dir
* xlog_test.go
refact code

## 库
* dailywriter.go
```
	"fmt"
	"os"
	"time"
```
* xlog.go
```
	"fmt"
	"log"
	"os"
	"path"
```
* xlog_test.go
```
	"testing"
	"time"
```

## 对下层的依赖

## 对外暴露的方法
* xlog.XLog.Crit
* xlog.XLog.Debug
* xlog.XLog.Error
* xlog.XLog.Info
* xlog.XLog.Warn
* xlog.NewDailyLog
