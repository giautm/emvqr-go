package internal

import (
	"fmt"

	"golang.org/x/text/currency"
)

type CurrencyFormatter struct {
	Amount interface{}
	Unit   currency.Unit
}

func (a *CurrencyFormatter) Format(s fmt.State, verb rune) {
	if a.Amount == nil {
		return
	}

	if _, ok := s.Precision(); !ok {
		scale, _ := currency.Cash.Rounding(a.Unit)
		fmt.Fprintf(s, "%.*f", scale, a.Amount)
	} else {
		fmt.Fprint(s, a.Amount)
	}
}

type Currency struct {
	Amount interface{}
	Code   string
}

var _ Valuer = (*Currency)(nil)

func (c *Currency) Value() (string, error) {
	unit, err := currency.ParseISO(c.Code)
	if err != nil {
		return "", err
	}

	s := fmt.Sprint(&CurrencyFormatter{Amount: c.Amount, Unit: unit})
	if len(s) > MaxLength {
		return "", ErrDataTooLong
	}

	return s, nil
}
