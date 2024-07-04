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

type ITransaction interface {
	// Hash() common.Hash
	// Valid() error

	AttachSignature(common.Signature)
	Serialize() ([]byte, error)
	Signature() common.Signature
	GenerationHashSeed() []byte
}

type Transaction struct {
	ITransaction

	version  uint8                 // transaction version		(1 byte)
	network  network.Network       // network information		(1 byte)
	txType   types.TransactionType // transaction type			(2 bytes)
	fee      types.MaxFee          // transaction max fee		(8 bytes)
	deadline types.Deadline        // transaction deadline		(8 bytes)

	verifiableEntityHeaderReserved1 []byte // reserved value 	(4 bytes)
	entityBodyReserved1             []byte // reserved value 	(4 bytes)

	size types.TransactionSize // transaction size			 	(4 bytes)

	signature common.Signature // transaction signature			(64 bytes)
	signer    common.PublicKey // transaction signer publickey	(32 bytes)
}

func (transaction Transaction) serialize() ([]byte, error) {
	// serialize common transaciton attrs
	serializeData := append(transaction.size.Bytes(), transaction.verifiableEntityHeaderReserved1[:]...)
	serializeData = append(serializeData, transaction.signature[:]...)
	serializeData = append(serializeData, transaction.signer[:]...)
	serializeData = append(serializeData, transaction.entityBodyReserved1[:]...)
	serializeData = append(serializeData, transaction.version)
	serializeData = append(serializeData, byte(transaction.network.Type))
	serializeData = append(serializeData, transaction.txType.Bytes()...)
	serializeData = append(serializeData, transaction.fee.Bytes()...)
	serializeData = append(serializeData, transaction.deadline.Bytes()...)

	return serializeData, nil
}

func (transaction *Transaction) AttachSignature(signature common.Signature) {
	transaction.signature = signature
}

func (transaction Transaction) Signature() common.Signature {
	return transaction.signature
}

func (transaction Transaction) Size() types.TransactionSize {
	return transaction.size
}

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

	appendedData := append(transactionFactory.network.GenerationHashSeed, data...)

	hasher := sha512.New()
	hasher.Write(appendedData)
	hashedData := hasher.Sum(nil)

	edPrivateKey := ed25519.NewKeyFromSeed(signer[:])
	sign, err := edPrivateKey.Sign(nil, hashedData,
		&ed25519.Options{
			Hash: crypto.SHA512,
		},
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
