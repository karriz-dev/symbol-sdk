package errors

import "errors"

var (
	ErrTxTypeNotFound = errors.New("tx type not found")
	ErrTxSerialize    = errors.New("tx serialize failed")
)
