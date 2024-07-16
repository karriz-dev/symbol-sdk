package factory

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha512"
	"time"

	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/signature"
	"github.com/karriz-dev/symbol-sdk/model/tx"
	"github.com/karriz-dev/symbol-sdk/network"
)

type TransactionFactory struct {
	signer   account.PublicKey
	network  network.Network
	maxFee   decimal.UInt64
	deadline decimal.UInt64
}

func NewTransactionFactory(network network.Network) *TransactionFactory {
	return &TransactionFactory{
		network:  network,
		maxFee:   decimal.NewUInt64(0),
		deadline: decimal.NewUInt64(0),
	}
}

func (transactionFactory *TransactionFactory) Signer(signerPublicKey account.PublicKey) *TransactionFactory {
	transactionFactory.signer = signerPublicKey

	return transactionFactory
}

func (transactionFactory *TransactionFactory) MaxFee(maxFee uint64) *TransactionFactory {
	transactionFactory.maxFee.Add(maxFee)

	return transactionFactory
}

func (transactionFactory *TransactionFactory) Deadline(deadline time.Duration) *TransactionFactory {
	transactionFactory.deadline.Add(transactionFactory.network.AddTime(deadline))

	return transactionFactory
}

func (transactionFactory TransactionFactory) Sign(tx tx.Transaction, signer account.PrivateKey) (signature.Signature, error) {
	// generation hash seed + except tx header data (common tx header length: 108)
	appendedData := append(transactionFactory.network.GenerationHashSeed, tx.Serialize()[108:]...)

	edPrivateKey := ed25519.NewKeyFromSeed(signer[:])
	sign, err := edPrivateKey.Sign(nil, appendedData,
		&ed25519.Options{},
	)
	if err != nil {
		return signature.Signature{}, err
	}

	return signature.Signature(sign), nil
}

func (transactionFactory TransactionFactory) Verify(payload []byte, signature []byte, signer account.PublicKey) error {
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

func (transactionFactory TransactionFactory) TransferTransactionV1(isEmbedded bool) tx.TransferTransactionV1 {
	// new transfer_transaction_v1
	return tx.NewTransferTransactionV1(
		transactionFactory.network,
		transactionFactory.maxFee,
		transactionFactory.deadline,
		transactionFactory.signer,
		isEmbedded,
	)
}
