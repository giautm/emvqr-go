# SGQR payload builder written in Go

## VietQR specification

### Data Objects Under the Root of a QR Code

| Name | ID | Format | Length | Presence | Comment |
| --- | --- | --- | --- | --- | --- |
| Payload Format Indicator | "00" | N | "02" | M | Refer to [Ref A].
| Point of Initiation Method | "01" | N | "02" | O | "11" for static, "22" for dynamic<br/>Refer to [Ref A].
| Merchant Account Information | "02"-"51" | ans | Each var. up to "99" | M | At least one Merchant Account Information data object shall be present.<br/>Refer to [Ref A].
| Transaction Currency | "53" | N | "03" | M | Fixed to "704"<br/>Refer to [Ref A].
| Country Code | "58" | ans | "02" | M | Fixed to "VN"<br/>Refer to [Ref A].
| Additional Data Field Template | "62" | S | var. up to "99" | O | Refer to [Ref A].
| CRC | "63" | ans | "04" | M | Refer to [Ref A].

### Data Objects for Additional Data Field Template (ID "62")

| Name | ID | Format | Length | Presence |
| --- | --- | --- | --- | --- |
| Purpose of Transaction | "08" | ans | var. up to "25" | O

**VietQR - Merchant Account Information (ID "38")**

| Data Object | Input Characters | Remarks|
| --- | --- | --- |
| Merchant Account Information | "3856" | Floating ID "38". This ID is allocated for this QR only<br/>Refer to [Ref B].
| - Global Unique Identifier<br/>- Beneficiary Organization<br/>	* ACQ ID / BNB ID<br/>	*  Merchant ID / Consumer ID<br/> - Service Code | "0010A000000727"<br/>"012600069704150112113366668888"<br/><br/><br/>"0208QRIBFTTA" | Reversed domain<br/><br/>970415 - [NAPAS BIN](https://www.sbv.gov.vn/webcenter/ShowProperty?nodeId=/UCMServer/SBV399939//idcPrimaryFile&revision=latestreleased)<br/>113366668888<br/>QRIBFTTA / QRIBFTTC

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

## References
- [Ref A: EMV® QR Code Specification for Payment Systems (EMV® QRCPS) Merchant-Presented Mode][Ref A]
- [Ref B: TÀI LIỆU QUY ĐỊNH VỀ ĐỊNH DẠNG MÃ VIETQR TRONG DỊCH VỤ CHUYỂN NHANH NAPAS247][Ref B]

[Ref A]: https://www.emvco.com/emv-technologies/qrcodes/
[Ref B]: https://vietqr.net/portal-service/download/documents/QR_Format_T&C_v1.0_VN_092021.pdf
