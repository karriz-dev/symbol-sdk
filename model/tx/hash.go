package tx

import (
	"encoding/hex"
	"strings"
)

type Hash [32]byte

func (h Hash) Hex() string {
	return strings.ToUpper(hex.EncodeToString(h[:]))
}

func (h Hash) String() string {
	return strings.ToUpper(hex.EncodeToString(h[:]))
}
