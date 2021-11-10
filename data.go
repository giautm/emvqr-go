package sgqr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const MaxLength = 99

var ErrStringTooLong = errors.New("string too long")

type Valuer interface {
	Value() (string, error)
}

type Pair struct {
	ID   string
	Data Valuer
}

type String string

func NewString(id, value string) Pair {
	return Pair{ID: id, Data: String(value)}
}

func (s String) Value() (string, error) {
	if len(s) > MaxLength {
		return "", ErrStringTooLong
	}

	return string(s), nil
}

type Array []Pair

func NewArray(id string, values ...Pair) Pair {
	return Pair{ID: id, Data: Array(values)}
}

func (arr Array) Value() (string, error) {
	buf := &strings.Builder{}
	for _, i := range arr {
		s, err := i.Data.Value()
		if err != nil {
			return "", err
		}

		fmt.Fprintf(buf, "%s%02d%s", i.ID, len(s), s)
		if buf.Len() > MaxLength {
			return "", ErrStringTooLong
		}
	}

	return buf.String(), nil
}

type Float64 float64

func NewFloat64(id string, value float64) Pair {
	return Pair{ID: id, Data: Float64(value)}
}

func (a Float64) Value() (string, error) {
	s := strconv.FormatFloat(float64(a), 'f', 2, 64)
	if len(s) > MaxLength {
		return "", ErrStringTooLong
	}

	return s, nil
}

type UInt64 uint64

func NewUInt64(id string, value uint64) Pair {
	return Pair{ID: id, Data: Float64(value)}
}

func (a UInt64) Value() (string, error) {
	s := strconv.FormatUint(uint64(a), 10)
	if len(s) > MaxLength {
		return "", ErrStringTooLong
	}

	return s, nil
}
