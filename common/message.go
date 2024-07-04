package common

import "encoding/binary"

type Message string

type MessageLength uint16

func (message Message) Bytes() []byte {
	if len(message) <= 0 {
		return []byte{}
	}

	return append([]byte{0x00}, []byte(message)...)
}

func (messageLength MessageLength) Bytes() []byte {
	messageLengthBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(messageLengthBytes, uint16(messageLength))

	return messageLengthBytes
}
