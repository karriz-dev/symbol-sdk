package tx

import (
	"encoding/hex"
	"strings"
)

type Payload []byte

func (p Payload) Hex() string {
	return strings.ToUpper(hex.EncodeToString(p))
}

func (p Payload) String() string {
	return strings.ToUpper(hex.EncodeToString(p))
}
