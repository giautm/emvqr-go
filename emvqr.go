package emvqr

import (
	"fmt"

	"giautm.dev/emvqr/internal"
)

var ErrDataTooLong = internal.ErrDataTooLong

type (
	Pair   = internal.Pair
	Valuer = internal.Valuer
)

func Currency(id string, isoCode string, amount interface{}) Pair {
	return Pair{ID: id, Data: &internal.Currency{
		Amount: amount,
		Code:   isoCode,
	}}
}

func List(id string, values ...Pair) Pair {
	return Pair{ID: id, Data: internal.List(values)}
}

func String(id, value string) Pair {
	return Pair{ID: id, Data: internal.String(value)}
}

func Float64(id string, value float64) Pair {
	return Pair{ID: id, Data: internal.Float64(value)}
}

func Uint64(id string, value uint64) Pair {
	return Pair{ID: id, Data: internal.Uint64(value)}
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
	return List("26", info...)
}

func MerchantCategory(code string) Pair {
	return String("52", code)
}

func TransactionCurrency(code string) Pair {
	return String("53", code)
}

func TransactionAmount(isoCode string, amount interface{}) Pair {
	return Currency("54", isoCode, amount)
}

func TransactionAmountUint(amount uint64) Pair {
	return Uint64("54", amount)
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
	return List("62", data...)
}

func BuildPayload(root ...Pair) (string, error) {
	s, err := internal.ListRoot(root).Value()
	if err != nil {
		return "", err
	}

	s += "6304"
	return fmt.Sprintf("%s%04X", s, crcCCITTFalse([]byte(s))), nil
}

func CheckCRC(input string) bool {
	l := len(input)
	if l <= 4 {
		return false
	}

	crc := fmt.Sprintf("%04X", crcCCITTFalse([]byte(input[:l-4])))
	return crc == input[l-4:]
}

func GetIn(input string, ids ...string) string {
	for _, id := range ids {
		found := false

		for remain := len(input); remain > 4; remain = len(input) {
			found = input[:2] == id[:2]
			last := 4 + int((input[2]-'0')*10+input[3]-'0')
			if last > remain {
				return "" // EOF
			}
			if found {
				input = input[4:last]
				break
			}

			input = input[last:]
		}

		if !found && len(input) <= 4 {
			return "" // EOF
		}
	}

	return input
}
