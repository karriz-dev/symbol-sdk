package model

import (
	"errors"
)

var (
	ErrInvalidRef      = errors.New("invalid output ref")
	ErrInvalidSize     = errors.New("invalid data size")
	ErrUnknownBaseType = errors.New("unknown base type")
)

type Base interface {
	String() string // plain string value
	Hex() string    // bytes to hex string
	Bytes() []byte  // base value to bytes
}
