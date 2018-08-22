# ecli
## 文件内容及功能描述
* main.go
删除node相关的子命令，新增发行token的命令

## 库
* main.go
```
	"github.com/EducationEKT/EKT/cmd/ecli/cmd"
	"github.com/EducationEKT/EKT/cmd/ecli/param"
	"github.com/spf13/cobra"
```

## 对下层的依赖
* cmd.TransactionCmd
* cmd.AccountCmd
* cmd.NodeCommand
* cobra.Command
## 对外暴露的方法

