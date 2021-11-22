package sgqr

type VietQRInput struct {
	AcqID          string `json:"acqID"`
	AccountNo      string `json:"accountNo"`
	Amount         uint64 `json:"amount"`
	AdditionalInfo string `json:"addInfo"`
}

func (i VietQRInput) BuildPayload() (string, error) {
	pairs := []Pair{
		PayloadFormatIndicator(),
		PointOfInitiationMethod(i.Amount > 0),
		List("38",
			String("00", "A000000727"),
			List("01",
				String("00", i.AcqID),
				String("01", i.AccountNo),
			),
			// QR Inter-Bank Funds Transfer To Account
			String("02", "QRIBFTTA"),
		),
		TransactionCurrency("704"),
		CountryCode("VN"),
	}
	if i.Amount > 0 {
		pairs = append(pairs, TransactionAmountUint(i.Amount))
	}
	if len(i.AdditionalInfo) > 0 {
		pairs = append(pairs, AdditionalData(String("08", i.AdditionalInfo)))
	}

	return BuildPayload(pairs...)
}
