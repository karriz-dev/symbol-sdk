package account

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/karriz-dev/symbol-sdk/network"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	ErrInvalidKeyLength = errors.New("invalid key length")
)

type PrivateKey [32]byte
type PublicKey [32]byte

func PrivateKeyFromHex(hexString string) (PrivateKey, error) {
	decodedPrivateKey, err := hex.DecodeString(hexString)
	if err != nil {
		return PrivateKey{}, err
	}
	if len(decodedPrivateKey) != 32 {
		return PrivateKey{}, ErrInvalidKeyLength
	}

	return PrivateKey([]byte(ed25519.NewKeyFromSeed(decodedPrivateKey)[:32])), nil
}

func (privateKey PrivateKey) String() string {
	return strings.ToUpper(hex.EncodeToString(privateKey[:]))
}

func (privateKey PrivateKey) Hex() string {
	return strings.ToUpper(hex.EncodeToString(privateKey[:]))
}

func PublicKeyFromHex(hexString string) (PublicKey, error) {
	decodedPrivateKey, err := hex.DecodeString(hexString)
	if err != nil {
		return PublicKey{}, err
	}
	if len(decodedPrivateKey) != 32 {
		return PublicKey{}, ErrInvalidKeyLength
	}

	return PublicKey([]byte(ed25519.NewKeyFromSeed(decodedPrivateKey)[32:])), nil
}

func (publicKey PublicKey) String() string {
	return strings.ToUpper(hex.EncodeToString(publicKey[:]))
}

func (publicKey PublicKey) Hex() string {
	return strings.ToUpper(hex.EncodeToString(publicKey[:]))
}

func (publicKey PublicKey) Address(network network.Network) Address {
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
