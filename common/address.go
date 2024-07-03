package common

import (
	"crypto/ed25519"
	"encoding/base32"
	"encoding/hex"
	"strings"

	"github.com/karriz-dev/symbol-sdk/network"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type PrivateKey [32]byte
type PublicKey [32]byte

type Address [24]byte

type KeyPair struct {
	PrivateKey PrivateKey
	PublicKey  PublicKey
}

func NewKeyPair() (KeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		PrivateKey: PrivateKey(privateKey[:32]),
		PublicKey:  PublicKey(publicKey),
	}, nil
}

func HexToPrivateKey(privateKeyHex string) (PrivateKey, error) {
	hextoBytes, err := HexToBytes(privateKeyHex)
	if err != nil {
		return PrivateKey{}, nil
	}

	privateKey := ed25519.NewKeyFromSeed(hextoBytes)

	return PrivateKey(privateKey), nil
}

func HexToAddress(addressHex string) (Address, error) {
	hextoBytes, err := HexToBytes(addressHex)
	if err != nil {
		return Address{}, nil
	}

	return Address(hextoBytes), nil
}

func HexToKeyPair(privateKeyHex string) (KeyPair, error) {
	hextoBytes, err := HexToBytes(privateKeyHex)
	if err != nil {
		return KeyPair{}, nil
	}

	privateKey := ed25519.NewKeyFromSeed(hextoBytes)
	publicKey := privateKey.Public().(ed25519.PublicKey)

	return KeyPair{
		PrivateKey: PrivateKey(privateKey),
		PublicKey:  PublicKey(publicKey),
	}, nil
}

func DecodeAddress(encodedAddress string) (Address, error) {
	base32Decoder := base32.StdEncoding.WithPadding(base32.NoPadding)

	decodeAddr, err := base32Decoder.DecodeString(encodedAddress)
	if err != nil {
		return Address{}, nil
	}

	return Address(decodeAddr), nil
}

func PublicKeyToAddress(publicKey PublicKey, network network.Network) Address {
	// step 1. sha3_256 hash of the publickey
	sha3Hasher := sha3.New256()
	sha3Hasher.Write(publicKey[:])
	publicKeyHash := sha3Hasher.Sum(nil)

	// step 2. ripemd160 hash of (1)
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(publicKeyHash)
	addressHash := ripemd160Hasher.Sum(nil)

	// step 3. add network identifier byte in front of (2)
	addressWithNetworkId := make([]byte, 1)
	addressWithNetworkId[0] = byte(network.Type)
	addressWithNetworkId = append(addressWithNetworkId, addressHash...)

	// step 4. sha3_256 for checksum
	sha3FinalHasher := sha3.New256()
	sha3FinalHasher.Write(addressWithNetworkId)
	finalHash := sha3FinalHasher.Sum(nil)

	// step 5. return (3) + (4)[:3] (in catapult's address checksum size = 3)
	return Address(append(addressWithNetworkId, finalHash[:3]...))
}

func (privateKey PrivateKey) Hex() string {
	encodedHex := hex.EncodeToString(privateKey[:])

	return strings.ToUpper(encodedHex)
}

func (publicKey PublicKey) Hex() string {
	encodedHex := hex.EncodeToString(publicKey[:])

	return strings.ToUpper(encodedHex)
}

func (address Address) Encode() string {
	base32Encoder := base32.StdEncoding.WithPadding(base32.NoPadding)

	return base32Encoder.EncodeToString(address[:])
}
