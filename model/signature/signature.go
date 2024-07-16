package signature

import (
	"encoding/hex"
	"strings"
)

type Signature [64]byte

func (signature *Signature) FromHex() string {
	return strings.ToUpper(hex.EncodeToString(signature[:]))
}

func FromHex(hexString string) Signature {
	decodedSignature, err := hex.DecodeString(hexString)
	if err != nil {
		return Signature{}
	}
	if len(decodedSignature) != 64 {
		return Signature{}
	}

	return Signature(decodedSignature)
}

func (signature Signature) Hex() string {
	return strings.ToUpper(hex.EncodeToString(signature[:]))
}

func (signature Signature) String() string {
	return strings.ToUpper(hex.EncodeToString(signature[:]))
}
