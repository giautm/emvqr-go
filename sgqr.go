package sgqr

import (
	"fmt"
)

func PayloadFormatIndicator() Pair {
	return NewString("00", "01")
}

func PointOfInitiationMethod(isDynamic bool) Pair {
	val := "11"
	if isDynamic {
		val = "12"
	}

	return NewString("01", val)
}

func MerchantAccountInfo(info ...Pair) Pair {
	return NewArray("26", info...)
}

func MerchantCategory(code string) Pair {
	return NewString("52", code)
}

func TransactionCurrency(code string) Pair {
	return NewString("53", code)
}

func TransactionAmount(amount float64) Pair {
	return NewFloat64("54", amount)
}

func CountryCode(code string) Pair {
	return NewString("58", code)
}

func MerchantName(name string) Pair {
	return NewString("59", name)
}

func MerchantCity(city string) Pair {
	return NewString("60", city)
}

func PostalCode(code string) Pair {
	return NewString("61", code)
}

func AdditionalData(data ...Pair) Pair {
	return NewArray("62", data...)
}

func BuildPayload(root Array) (string, error) {
	s, err := root.Value()
	if err != nil {
		return "", err
	}

	s += "6304"
	return fmt.Sprintf("%s%04X", s, crcCCITTFalse([]byte(s))), nil
}
