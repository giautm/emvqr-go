package emvqr

const (
	GUIDVietQR = "A000000727"
)

type ServiceCode string

func (s ServiceCode) ToPair(id string) Pair {
	return String(id, string(ServiceCodeToAccount))
}

const (
	// QR Inter-Bank Funds Transfer To Account
	ServiceCodeToAccount ServiceCode = "QRIBFTTA"
	// QR Inter-Bank Funds Transfer To Card
	ServiceCodeToCard ServiceCode = "QRIBFTTC"
)

type VietQRInput struct {
	AcqID          string `json:"acqID"`
	AccountNo      string `json:"accountNo"`
	Amount         uint64 `json:"amount"`
	AdditionalInfo string `json:"addInfo"`
	IsCard         bool   `json:"isCard"`
}

func (i VietQRInput) BuildPayload() (string, error) {
	serviceCode := ServiceCodeToAccount
	if i.IsCard {
		serviceCode = ServiceCodeToCard
	}

	pairs := []Pair{
		PayloadFormatIndicator(),
		PointOfInitiationMethod(i.Amount > 0),
		List("38",
			String("00", GUIDVietQR),
			List("01",
				String("00", i.AcqID),
				String("01", i.AccountNo),
			),
			serviceCode.ToPair("02"),
		),
		TransactionCurrency("704"),
		CountryCode("VN"),
	}
	if i.Amount > 0 {
		// VND has the precision of 0 digits, so we need to
		// use TransactionAmountUint instead of
		// TransactionAmount to optimize performance.
		pairs = append(pairs, TransactionAmountUint(i.Amount))
	}
	if len(i.AdditionalInfo) > 0 {
		pairs = append(pairs, AdditionalData(String("08", i.AdditionalInfo)))
	}

	return BuildPayload(pairs...)
}
