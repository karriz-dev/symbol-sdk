package common

import "encoding/hex"

func HexToBytes(hexString string) ([]byte, error) {
	result, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return result, nil
}
