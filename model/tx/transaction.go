package tx

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/signature"
	"github.com/karriz-dev/symbol-sdk/network"
)

type Transaction interface {
	Serialize() ([]byte, error)
	Size() uint32
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
	hash       Hash // tx hash
}

func (tx BaseTransaction) Serialize() ([]byte, error) {
	// TODO:: check error case
	var serializeData []byte

	if tx.isEmbedded {
		tx.size.Add(48) // embedded tx header size: 48 bytes

		serializeData = append(tx.size.Bytes(), tx.verifiableEntityHeaderReserved1.Bytes()...)
		serializeData = append(serializeData, tx.signer[:]...)
		serializeData = append(serializeData, tx.entityBodyReserved1.Bytes()...)
		serializeData = append(serializeData, tx.version.Bytes()...)
		serializeData = append(serializeData, byte(tx.network.Type))
		serializeData = append(serializeData, tx.txType.Bytes()...)
	} else {
		tx.size.Add(128) // base tx header size: 128 bytes

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
