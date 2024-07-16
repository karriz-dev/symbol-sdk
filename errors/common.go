package errors

import "errors"

var (
	ErrGetBytes    = errors.New("unkown type")
	ErrSizeInvalid = errors.New("convert type size not matched")
)
