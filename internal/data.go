package internal

import (
	"errors"
	"strconv"
	"strings"
)

const MaxLength = 99

var (
	ErrDataTooLong = errors.New("emvqr: data too long")
	digits         = []rune("0123456789")
)

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

type list struct {
	Pairs []Pair
	root  bool
}

func List(pairs []Pair) Valuer {
	return list{Pairs: pairs, root: false}
}

func ListRoot(pairs []Pair) Valuer {
	return list{Pairs: pairs, root: true}
}

func (arr list) Value() (string, error) {
	buf := &strings.Builder{}
	buf.Grow(96)
	for _, i := range arr.Pairs {
		s, err := i.Data.Value()
		if err != nil {
			return "", err
		}

		l := len(s)
		buf.WriteString(i.ID[:2])
		buf.WriteRune(digits[l/10])
		buf.WriteRune(digits[l%10])
		buf.WriteString(s)
		if !arr.root && buf.Len() > MaxLength {
			return "", ErrDataTooLong
		}
	}

	return buf.String(), nil
}
