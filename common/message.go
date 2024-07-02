package common

type Message string

type MessageLength uint16

func (message Message) Bytes() []byte {
	if len(message) <= 0 {
		return []byte{}
	}

	return append([]byte{0x00}, []byte(message)...)
}

func (messageLength MessageLength) Bytes() []byte {
	return []byte{byte(messageLength >> 8), byte(messageLength & 0xff)}
}
