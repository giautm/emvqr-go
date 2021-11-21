# SGQR payload builder written in Go

## VietQR specification

**Document**

**Sample code**

```go
package main

import (
	"fmt"

	"giautm.dev/sgqr"
	qrcode2 "github.com/lizebang/qrcode-terminal"
	"github.com/skip2/go-qrcode"
)

func main() {
	i := &sgqr.VietQRInput{
		AcqID:          "970415",
		AccountNo:      "113366668888",
		Amount:         79000,
		AdditionalInfo: "Ung Ho Quy Vac Xin",
	}

	payload, err := i.BuildPayload()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data", payload)

	qrcode2.QRCode(payload, qrcode2.BrightWhite, qrcode2.NormalBlack, qrcode.Medium)
}
```

Performance
```
goos: darwin
goarch: amd64
pkg: giautm.dev/sgqr
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
Benchmark_BuildPayload
Benchmark_BuildPayload-8   	 1558408	       771.8 ns/op	     594 B/op	       8 allocs/op
PASS
ok  	giautm.dev/sgqr	2.447s
```