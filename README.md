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
	i := &VietQRInput{
		AcqID:     "970415",
		AccountNo: "113366668888",
		Amount:    79000,
		Message:   "Ung Ho Quy Vac Xin",
	}

	payload, err := i.BuildPayload()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data", payload)

	qrcode2.QRCode(payload, qrcode2.BrightWhite, qrcode2.NormalBlack, qrcode.Medium)
}

type VietQRInput struct {
	AcqID     string `json:"acqID"`
	AccountNo string `json:"accountNo"`
	Amount    uint64 `json:"amount"`
	Message   string `json:"message"`
}

func (i VietQRInput) BuildPayload() (string, error) {
	pairs := []sgqr.Pair{
		sgqr.PayloadFormatIndicator(),
		sgqr.PointOfInitiationMethod(i.Amount > 0),
		sgqr.List("38",
			sgqr.String("00", "A000000727"),
			sgqr.List("01",
				sgqr.String("00", i.AcqID),
				sgqr.String("01", i.AccountNo),
			),
			// QR Inter-Bank Funds Transfer
			sgqr.String("02", "QRIBFTTA"),
		),
		sgqr.TransactionAmountUint(i.Amount),
		sgqr.TransactionCurrency("704"),
		sgqr.CountryCode("VN"),
	}
	if len(i.Message) > 0 {
		pairs = append(pairs, sgqr.AdditionalData(sgqr.String("08", i.Message)))
	}

	return sgqr.BuildPayload(pairs...)
}
```