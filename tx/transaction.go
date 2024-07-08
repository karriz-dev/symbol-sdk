package tx

import (
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/karriz-dev/symbol-sdk/types"
)

type ITransaction interface {
	AttachSignature(common.Signature)
	Serialize() ([]byte, error)
	Signature() common.Signature
}

type IEmbeddedTransaction interface {
	Serialize() ([]byte, error)
}

type Transaction struct {
	ITransaction

	size                            types.TransactionSize // transaction size			(4 bytes)
	verifiableEntityHeaderReserved1 []byte                // reserved value 			(4 bytes)
	signature                       common.Signature      // transaction signature		(64 bytes)

	// Entity Body
	signer              common.PublicKey      // transaction signer publickey	(32 bytes)
	entityBodyReserved1 []byte                // reserved value 				(4 bytes)
	version             uint8                 // transaction version			(1 byte)
	network             network.Network       // network information			(1 byte)
	txType              types.TransactionType // transaction type				(2 bytes)
	fee                 types.MaxFee          // transaction max fee			(8 bytes)
	deadline            types.Deadline        // transaction deadline			(8 bytes)
}

type EmbeddedTransaction struct {
	IEmbeddedTransaction

	size                               types.TransactionSize // transaction size			 	(4 bytes)
	embeddedTransactionHeaderReserved1 []byte                // reserved value 					(4 bytes)
	signer                             common.PublicKey      // transaction signer publickey	(32 bytes)
	entityBodyReserved1                []byte                // reserved value 					(4 bytes)

	// Entity Body
	version uint8                 // transaction version				(1 byte)
	network network.Network       // network information				(1 byte)
	txType  types.TransactionType // transaction type					(2 bytes)
}

func (transactionHeader *Transaction) AttachSignature(signature common.Signature) {
	transactionHeader.signature = signature
}

func (transactionHeader Transaction) serialize() ([]byte, error) {
	serializeData := append(transactionHeader.size.Bytes(), transactionHeader.verifiableEntityHeaderReserved1[:]...)
	serializeData = append(serializeData, transactionHeader.signature[:]...)
	serializeData = append(serializeData, transactionHeader.signer[:]...)
	serializeData = append(serializeData, transactionHeader.entityBodyReserved1[:]...)
	serializeData = append(serializeData, transactionHeader.version)
	serializeData = append(serializeData, byte(transactionHeader.network.Type))
	serializeData = append(serializeData, transactionHeader.txType.Bytes()...)
	serializeData = append(serializeData, transactionHeader.fee.Bytes()...)
	serializeData = append(serializeData, transactionHeader.deadline.Bytes()...)

	return serializeData, nil
}

func (transactionHeader Transaction) Signature() common.Signature {
	return transactionHeader.signature
}

func (embeddedTransactionHeader EmbeddedTransaction) serialize() ([]byte, error) {
	serializeData := append(embeddedTransactionHeader.size.Bytes(), embeddedTransactionHeader.embeddedTransactionHeaderReserved1[:]...)
	serializeData = append(serializeData, embeddedTransactionHeader.signer[:]...)
	serializeData = append(serializeData, embeddedTransactionHeader.entityBodyReserved1[:]...)
	serializeData = append(serializeData, embeddedTransactionHeader.version)
	serializeData = append(serializeData, byte(embeddedTransactionHeader.network.Type))
	serializeData = append(serializeData, embeddedTransactionHeader.txType.Bytes()...)

	return serializeData, nil
}
