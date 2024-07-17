package tx

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/signature"
	"github.com/karriz-dev/symbol-sdk/network"
)

const (
	BaseTransactionSize     = 48
	EmbeddedTransactionSize = 128

	TransactionHeaderSize = 108
	AggregateHashedSize   = 52
)

type Transaction interface {
	Serialize() ([]byte, error)
	Size() uint32
	Payload() (Payload, error)
}

type BaseTransaction struct {
	Transaction

	size                            decimal.UInt32      // transaction size			(4 bytes)
	verifiableEntityHeaderReserved1 decimal.UInt32      // reserved value 				(4 bytes)
	signature                       signature.Signature // transaction signature		(64 bytes)

	signer              account.PublicKey // transaction signer publickey	(32 bytes)
	entityBodyReserved1 decimal.UInt32    // reserved value 				(4 bytes)
	version             decimal.UInt8     // transaction version			(1 byte)
	network             network.Network   // network information			(1 byte)
	txType              decimal.UInt16    // transaction type				(2 bytes)
	fee                 decimal.UInt64    // transaction max fee			(8 bytes)
	deadline            decimal.UInt64    // transaction deadline			(8 bytes)

	isEmbedded bool // check embedded tx
}

func (tx *BaseTransaction) SetBaseSize(innerTxSize uint32, isEmbedded bool) {
	if isEmbedded {
		tx.size = decimal.NewUInt32(innerTxSize + 48)
	} else {
		tx.size = decimal.NewUInt32(innerTxSize + 128)
	}
}

func (tx BaseTransaction) Serialize() ([]byte, error) {
	// TODO:: check error case
	var serializeData []byte

	if tx.isEmbedded {
		serializeData = append(tx.size.Bytes(), tx.verifiableEntityHeaderReserved1.Bytes()...)
		serializeData = append(serializeData, tx.signer[:]...)
		serializeData = append(serializeData, tx.entityBodyReserved1.Bytes()...)
		serializeData = append(serializeData, tx.version.Bytes()...)
		serializeData = append(serializeData, byte(tx.network.Type))
		serializeData = append(serializeData, tx.txType.Bytes()...)
	} else {
		serializeData = append(tx.size.Bytes(), tx.verifiableEntityHeaderReserved1.Bytes()...)
		serializeData = append(serializeData, tx.signature[:]...)
		serializeData = append(serializeData, tx.signer[:]...)
		serializeData = append(serializeData, tx.entityBodyReserved1.Bytes()...)
		serializeData = append(serializeData, tx.version.Bytes()...)
		serializeData = append(serializeData, byte(tx.network.Type))
		serializeData = append(serializeData, tx.txType.Bytes()...)
		serializeData = append(serializeData, tx.fee.Bytes()...)
		serializeData = append(serializeData, tx.deadline.Bytes()...)
	}

	return serializeData, nil
}

func (tx BaseTransaction) Size() uint32 {
	return tx.size.Value()
}

func (tx *BaseTransaction) AttachSignature(signature signature.Signature) {
	tx.signature = signature
}
