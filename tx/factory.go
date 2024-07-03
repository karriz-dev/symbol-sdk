package tx

import (
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/karriz-dev/symbol-sdk/types"
)

type ITransaction interface {
	// Hash() common.Hash
	// Valid() error

	Serialize() ([]byte, error)
}

type Transaction struct {
	ITransaction

	version  uint8                 // transaction version
	network  network.Network       // network information
	txType   types.TransactionType // transaction type
	fee      types.MaxFee          // transaction max fee
	deadline types.Deadline        // transaction deadline

	verifiableEntityHeaderReserved1 []byte // reserved value 4 bytes
	entityBodyReserved1             []byte // reserved value 4 bytes

	size types.TransactionSize // transaction size

	signature common.Signature // transaction signature
	signer    common.KeyPair   // transaction signer publickey
}

func (transaction Transaction) Serialize() ([]byte, error) {
	// serialize common transaciton attrs
	serializeData := append(transaction.size.Bytes(), transaction.verifiableEntityHeaderReserved1[:]...)
	serializeData = append(serializeData, transaction.signature[:]...)
	serializeData = append(serializeData, transaction.signer.PublicKey[:]...)
	serializeData = append(serializeData, transaction.entityBodyReserved1[:]...)
	serializeData = append(serializeData, transaction.version)
	serializeData = append(serializeData, byte(transaction.network.Type))
	serializeData = append(serializeData, transaction.txType.Bytes()...)
	serializeData = append(serializeData, transaction.fee.Bytes()...)
	serializeData = append(serializeData, transaction.deadline.Bytes()...)

	return serializeData, nil
}

func (transaction Transaction) Sign() error {
	return nil
}

type TransactionFactory struct {
	signer   common.KeyPair
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

func (transactionFactory *TransactionFactory) Signer(signerKeyPair common.KeyPair) *TransactionFactory {
	transactionFactory.signer = signerKeyPair

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
