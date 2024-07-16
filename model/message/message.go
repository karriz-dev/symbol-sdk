package message

import (
	"encoding/binary"
)

type Message string

func (message Message) Bytes() []byte {
	if len(message) <= 0 {
		return []byte{}
	}

	return append([]byte{0x00}, []byte(message)...)
}

func (message Message) LenBytes() []byte {
	messageLength := uint16(0)

	if len(message) <= 0 {
		return binary.LittleEndian.AppendUint16(nil, messageLength)
	}

	messageLength += 1

	return binary.LittleEndian.AppendUint16(nil, messageLength)
}
