# enode
## 文件内容及功能描述
* main.go
delete unused log
## 库
* main.go
```
	"flag"
	"fmt"
	"net/http"
	"os"
	_ "github.com/EducationEKT/EKT/api"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/db"
	"github.com/EducationEKT/EKT/log"
	"github.com/EducationEKT/EKT/param"
	"runtime"
	"github.com/EducationEKT/xserver/x_http"
```
## 对下层的依赖
* blockchain_manager.Init
* conf.InitConfig
* db.InitEKTDB
* log.InitLog
* param.InitBootNodes

## 对外暴露的方法

