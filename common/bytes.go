package common

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"strings"

	"golang.org/x/crypto/sha3"
)

func aa() {
	
}



// func From[T ~uint | []byte](data T) []byte {

// 	switch v := data.(type) {
// 	case uint16:
// 		result := make([]byte, 2)
// 		binary.LittleEndian.PutUint16(result, v)
// 		return result
// 	case uint:
// 		result := make([]byte, 4)
// 		binary.LittleEndian.PutUint32(result, uint32(v))
// 		return result
// 	case uint32:
// 		result := make([]byte, 4)
// 		binary.LittleEndian.PutUint32(result, uint32(v))
// 		return result
// 	case uint64:
// 		result := make([]byte, 8)
// 		binary.LittleEndian.PutUint64(result, v)
// 		return result
// 	default:
// 		return []byte{}
// 	}
// }

// func (bytes Bytes) Hex() string {
// 	encodedString := hex.EncodeToString(bytes)

// 	return strings.ToUpper(encodedString)
// }

// func BytesToHex(data []byte) string {
// 	encodedString := hex.EncodeToString(data)

// 	return strings.ToUpper(encodedString)
// }

// func BytesToHash(data []byte) (Hash, error) {
// 	hasher := sha3.New256()

// 	_, err := hasher.Write(data)
// 	if err != nil {
// 		return Hash{}, err
// 	}

// 	return Hash(hasher.Sum(nil)), nil
// }

// func BytesToJSONPayload(payload []byte) string {
// 	return "{\"payload\":\"" + BytesToHex(payload) + "\"}"
// }
