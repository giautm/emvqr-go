package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const MaxLength = 99

var ErrDataTooLong = errors.New("sgqr: data too long")

type Valuer interface {
	Value() (string, error)
}

type Pair struct {
	ID   string
	Data Valuer
}

type String string

func (s String) Value() (string, error) {
	if len(s) > MaxLength {
		return "", ErrDataTooLong
	}

	return string(s), nil
}

type Array struct {
	Pairs []Pair
	Root  bool
}

func (arr Array) Value() (string, error) {
	buf := &strings.Builder{}
	for _, i := range arr.Pairs {
		s, err := i.Data.Value()
		if err != nil {
			return "", err
		}

		fmt.Fprintf(buf, "%s%02d%s", i.ID, len(s), s)
		if !arr.Root && buf.Len() > MaxLength {
			return "", ErrDataTooLong
		}
	}

	return buf.String(), nil
}

type Float64 float64

func (a Float64) Value() (string, error) {
	s := strconv.FormatFloat(float64(a), 'f', 2, 64)
	if len(s) > MaxLength {
		return "", ErrDataTooLong
	}

	return s, nil
}

type Uint64 uint64

func (a Uint64) Value() (string, error) {
	s := strconv.FormatUint(uint64(a), 10)
	if len(s) > MaxLength {
		return "", ErrDataTooLong
	}

	return s, nil
}
