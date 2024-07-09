package tx

import (
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/karriz-dev/symbol-sdk/types"
)

type ITransaction interface {
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

	isEmbedded bool // check embedded tx
}

func (tx Transaction) serialize() ([]byte, error) {
	if tx.isEmbedded {
		// embedded tx default size: 48 bytes
		tx.size += 48
	} else {
		// tx default size: 128 bytes
		tx.size += 128
	}

	serializeData := append(tx.size.Bytes(), tx.verifiableEntityHeaderReserved1[:]...)
	serializeData = append(serializeData, tx.signature[:]...)
	serializeData = append(serializeData, tx.signer[:]...)
	serializeData = append(serializeData, tx.entityBodyReserved1[:]...)
	serializeData = append(serializeData, tx.version)
	serializeData = append(serializeData, byte(tx.network.Type))
	serializeData = append(serializeData, tx.txType.Bytes()...)
	serializeData = append(serializeData, tx.fee.Bytes()...)
	serializeData = append(serializeData, tx.deadline.Bytes()...)

	return serializeData, nil
}

func (tx *Transaction) AttachSignature(signature common.Signature) {
	tx.signature = signature
}

func (tx Transaction) Signature() common.Signature {
	return tx.signature
}
