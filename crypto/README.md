# crypto
## 文件内容及功能描述
* crypto.go
refact code
* crypto_test.go
refact code
* sha3.go
refact code
* sha3_test.go
refact code

## 库
* crypto.go
```
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
```
* crypto_test.go
```
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
```
* sha3.go
```
	"bytes"
	"errors"
	"github.com/ethereum/go-ethereum/crypto/sha3"
```
* sha3_test.go
```
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
```

## 对下层的依赖

## 对外暴露的方法
* crypto.RecoverPubKey
* crypto.Sha3_256
* crypto.Crypto
* crypto.PubKey
* crypto.GenerateKeyPair
