# util
## 文件内容及功能描述
* http.go
bug fixed: 当url不通的时候，会导致panic
* type_cvt.go
重构round相关的逻辑，在round中不再处理dpos相关逻辑

## 库
* http.go
```
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
```
* type_cvt.go
```
	"bytes"
	"encoding/binary"
```

## 对下层的依赖

## 对外暴露的方法
* util.BytesToInt
取随机数哈希22位之后的全部数字
* util.HttpGet
* util.HttpPost

