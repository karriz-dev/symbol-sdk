package tx

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha512"
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/karriz-dev/symbol-sdk/types"
)

type TransactionFactory struct {
	signer   common.PublicKey
	network  network.Network
	maxFee   types.MaxFee
	deadline types.Deadline
}

func NewTransactionFactory(network network.Network) *TransactionFactory {
	return &TransactionFactory{
		network:  network,
		maxFee:   0,
		deadline: 0,
	}
}

func (transactionFactory *TransactionFactory) Signer(signerPublicKey common.PublicKey) *TransactionFactory {
	transactionFactory.signer = signerPublicKey

	return transactionFactory
}

func (transactionFactory *TransactionFactory) MaxFee(maxFee uint64) *TransactionFactory {
	transactionFactory.maxFee = types.MaxFee(maxFee)

	return transactionFactory
}

func (transactionFactory *TransactionFactory) Deadline(deadline time.Duration) *TransactionFactory {
	transactionFactory.deadline = types.Deadline(transactionFactory.network.AddTime(deadline))

	return transactionFactory
}

func (transactionFactory TransactionFactory) Sign(transaction ITransaction, signer common.PrivateKey) (common.Signature, error) {
	data, err := transaction.Serialize()
	if err != nil {
		return common.Signature{}, err
	}

	// generation hash seed + except tx header data (common tx header length: 108)
	appendedData := append(transactionFactory.network.GenerationHashSeed, data[108:]...)

	edPrivateKey := ed25519.NewKeyFromSeed(signer[:])
	sign, err := edPrivateKey.Sign(nil, appendedData,
		&ed25519.Options{},
	)
	if err != nil {
		return common.Signature{}, err
	}

	return common.Signature(sign), nil
}

func (transactionFactory TransactionFactory) Verify(payload []byte, signature []byte, signer common.PublicKey) error {
	appendedData := append(transactionFactory.network.GenerationHashSeed, payload...)

	hasher := sha512.New()
	hasher.Write(appendedData)
	hashedData := hasher.Sum(nil)

	err := ed25519.VerifyWithOptions((ed25519.PublicKey)(signer[:]), hashedData, signature,
		&ed25519.Options{
			Hash: crypto.SHA512,
		},
	)

	return err
}
