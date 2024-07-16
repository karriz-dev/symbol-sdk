package account

import (
	"encoding/base32"
	"encoding/hex"
	"strings"
)

type Address [24]byte

func AddressFromEncode(base32Address string) Address {
	base32Decoder := base32.StdEncoding.WithPadding(base32.NoPadding)

	decodeAddr, err := base32Decoder.DecodeString(base32Address)
	if err != nil {
		return Address{}
	}

	return Address(decodeAddr)
}

func AddressFromHex(hexString string) (Address, error) {
	decodedAddress, err := hex.DecodeString(hexString)
	if err != nil {
		return Address{}, err
	}
	if len(decodedAddress) != 24 {
		return Address{}, ErrInvalidKeyLength
	}

	return Address(decodedAddress), nil
}

func (address Address) String() string {
	return address.EncodedAddress()
}

func (address Address) Hex() string {
	encodedHex := hex.EncodeToString(address[:])

	return strings.ToUpper(encodedHex)
}

func (address Address) EncodedAddress() string {
	base32Encoder := base32.StdEncoding.WithPadding(base32.NoPadding)

	return base32Encoder.EncodeToString(address[:])
}
