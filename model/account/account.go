package account

import (
	"crypto/ed25519"

	"github.com/karriz-dev/symbol-sdk/network"
)

type Account struct {
	PrivateKey
	PublicKey
	Address
}

func NewRandomAccount(network network.Network) (Account, error) {
	_, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return Account{}, err
	}

	publicKey := PublicKey(privateKey[32:])

	return Account{
		PrivateKey: PrivateKey(privateKey),
		PublicKey:  publicKey,
		Address:    publicKey.Address(network),
	}, nil
}

func AccountFromPrivateKey(pri PrivateKey, network network.Network) (Account, error) {
	privateKey := ed25519.NewKeyFromSeed(pri[:])

	publicKey := PublicKey(privateKey[32:])

	return Account{
		PrivateKey: PrivateKey(privateKey),
		PublicKey:  publicKey,
		Address:    publicKey.Address(network),
	}, nil
}
