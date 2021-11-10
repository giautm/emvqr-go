package sgqr

import (
	"fmt"

	"giautm.dev/sgqr/internal"
)

var ErrDataTooLong = internal.ErrDataTooLong

type Pair = internal.Pair

func Array(id string, values ...Pair) Pair {
	return Pair{ID: id, Data: internal.Array(values)}
}

func String(id, value string) Pair {
	return Pair{ID: id, Data: internal.String(value)}
}

func Float64(id string, value float64) Pair {
	return Pair{ID: id, Data: internal.Float64(value)}
}

func Uint64(id string, value uint64) Pair {
	return Pair{ID: id, Data: internal.Float64(value)}
}

func PayloadFormatIndicator() Pair {
	return String("00", "01")
}

func PointOfInitiationMethod(isDynamic bool) Pair {
	val := "11"
	if isDynamic {
		val = "12"
	}

	return String("01", val)
}

func MerchantAccountInfo(info ...Pair) Pair {
	return Array("26", info...)
}

func MerchantCategory(code string) Pair {
	return String("52", code)
}

func TransactionCurrency(code string) Pair {
	return String("53", code)
}

func TransactionAmount(amount float64) Pair {
	return Float64("54", amount)
}

func CountryCode(code string) Pair {
	return String("58", code)
}

func MerchantName(name string) Pair {
	return String("59", name)
}

func MerchantCity(city string) Pair {
	return String("60", city)
}

func PostalCode(code string) Pair {
	return String("61", code)
}

func AdditionalData(data ...Pair) Pair {
	return Array("62", data...)
}

func BuildPayload(root ...Pair) (string, error) {
	s, err := internal.Array(root).Value()
	if err != nil {
		return "", err
	}

	s += "6304"
	return fmt.Sprintf("%s%04X", s, crcCCITTFalse([]byte(s))), nil
}
