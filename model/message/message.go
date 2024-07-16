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
	if len(message) <= 0 {
		return binary.LittleEndian.AppendUint16(nil, 0)
	}

	return binary.LittleEndian.AppendUint16(nil, uint16(len(message)+1))
}
