package tx

import (
	"symbol-sdk/common"
	"symbol-sdk/types"
	"time"
)

type ITransaction interface {
	// Hash() common.Hash
	// Valid() error

	Serialize() ([]byte, error)
}

type Transaction struct {
	ITransaction

	version  uint8                 // transaction version
	network  types.NetworkType     // network id
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
	serializeData = append(serializeData, byte(transaction.network))
	serializeData = append(serializeData, transaction.txType.Bytes()...)
	serializeData = append(serializeData, transaction.fee.Bytes()...)
	serializeData = append(serializeData, transaction.deadline.Bytes()...)

	return serializeData, nil
}

func (transaction Transaction) Sign() error {
	return nil
}

type TransactionFactory struct {
	signer      common.KeyPair
	networkType types.NetworkType
	maxFee      types.MaxFee
	deadline    types.Deadline
}

func NewTransactionFactory(networkType types.NetworkType) *TransactionFactory {
	return &TransactionFactory{
		networkType: networkType,
		maxFee:      0,
		deadline:    0,
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
	transactionFactory.deadline = types.Deadline(time.Now().Add(deadline).Unix())

	return transactionFactory
}
