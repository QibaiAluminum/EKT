# log
## 文件内容及功能描述
* log.go
update log init

## 库
* log.go
```
	"sync"
	"github.com/EducationEKT/EKT/xlog"
```

## 对下层的依赖
* xlog.XLog.Crit
* xlog.XLog.Debug
* xlog.XLog.Error
* xlog.XLog.Info
* xlog.XLog.Warn
* xlog.NewDailyLog

## 对外暴露的方法
* log.Crit
* log.Debug
* log.Info
* log.InitLog

