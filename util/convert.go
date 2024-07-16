package util

import (
	"encoding/hex"
	"strings"
)

func BytesToJSONPayload(datas []byte) string {
	return "{\"payload\": \"" + strings.ToUpper(hex.EncodeToString(datas)) + "\"}"
}
