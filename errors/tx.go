package errors

import "errors"

var (
	ErrTxTypeNotFound = errors.New("tx type not found")
	ErrTxSerialize    = errors.New("tx serialize failed")
)

// merkle error list
var (
	ErrEmptyTransaction = errors.New("transaction list size cannot be zero")
)

// transfer_tx error list
var (
	ErrRecipientNotValid = errors.New("recipient not valid")
)
