package common

import (
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/sha3"
)

type Hash [32]byte
type Signature [64]byte

var (
	Empty4Bytes = []byte{0x00, 0x00, 0x00, 0x00}
)

func HexToBytes(hexString string) ([]byte, error) {
	result, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func BytesToHex(data []byte) string {
	encodedString := hex.EncodeToString(data)

	return strings.ToUpper(encodedString)
}

func BytesToHash(data []byte) (Hash, error) {
	hasher := sha3.New256()

	_, err := hasher.Write(data)
	if err != nil {
		return Hash{}, err
	}

	return Hash(hasher.Sum(nil)), nil
}

func (address Address) Bytes() []byte {
	return address[:]
}
