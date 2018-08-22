# round
## 文件内容及功能描述
* round.go
超出范围的索引的bug修复

## 库
* round.go
```
	"encoding/json"
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/p2p"
	"github.com/EducationEKT/EKT/util"
```

## 对下层的依赖
* util.BytesToInt
取随机数哈希22位之后的全部数字

## 对外暴露的方法
* round.Round

