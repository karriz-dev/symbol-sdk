package tx

import (
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
)

type HashLockTransactionV1 struct {
	Transaction

	mosaic   common.Mosaic // Locked mosaic
	duration time.Duration // Number of blocks for which a lock should be valid. The default maximum is 48h (See the maxHashLockDuration network property).
	hash     common.Hash   // Hash of the AggregateBondedTransaction to be confirmed before unlocking the mosaics.
}

type EmbeddedHashLockTransactionV1 struct {
	EmbeddedTransaction

	mosaic   common.Mosaic // Locked mosaic
	duration time.Duration // Number of blocks for which a lock should be valid. The default maximum is 48h (See the maxHashLockDuration network property).
	hash     common.Hash   // Hash of the AggregateBondedTransaction to be confirmed before unlocking the mosaics.
}

func (transactionFactory *TransactionFactory) HashLockTransactionV1() HashLockTransactionV1 {
	tx := Transaction{
		size:                            128,
		version:                         0x01,
		network:                         transactionFactory.network,
		txType:                          0x4148,
		fee:                             transactionFactory.maxFee,
		deadline:                        transactionFactory.deadline,
		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
		signer:                          transactionFactory.signer,
	}

	// adding TransferTransactionV1 entity size
	tx.size += 56

	return HashLockTransactionV1{
		Transaction: tx,
	}
}

func (embeddedTransactionFactory *EmbeddedTransactionFactory) EmbeddedHashLockTransactionV1() EmbeddedHashLockTransactionV1 {
	tx := EmbeddedTransaction{
		size:                48, // 48 = embedded tx header default size
		version:             0x01,
		network:             embeddedTransactionFactory.network,
		txType:              0x4154, // transfer_transaction_v1
		entityBodyReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		signer:              embeddedTransactionFactory.signer,
	}

	// adding TransferTransactionV1 entity size
	tx.size += 56

	return EmbeddedHashLockTransactionV1{
		EmbeddedTransaction: tx,
	}
}

// func (transferTransactionV1 *TransferTransactionV1) Recipient(recipient common.Address) *TransferTransactionV1 {
// 	transferTransactionV1.recipient = recipient

// 	return transferTransactionV1
// }

// func (transferTransactionV1 TransferTransactionV1) Serialize() ([]byte, error) {
// 	// serialize inner common tx attrs
// 	serializeData, err := transferTransactionV1.Transaction.serialize()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// serialize attrs
// 	serializeData = append(serializeData, transferTransactionV1.recipient.Bytes()...)
// 	serializeData = append(serializeData, transferTransactionV1.messageLength.Bytes()...)
// 	serializeData = append(serializeData, transferTransactionV1.mosaicsCount.Byte())
// 	serializeData = append(serializeData, transferTransactionV1.transferTransactionBodyReserved_1)
// 	serializeData = append(serializeData, transferTransactionV1.transferTransactionBodyReserved_2[:]...)

// 	// serialize mosiacs
// 	for _, mosaic := range transferTransactionV1.mosaics {
// 		serializeData = append(serializeData, mosaic.Bytes()...)
// 	}

// 	// serialize message
// 	serializeData = append(serializeData, transferTransactionV1.message.Bytes()...)

// 	return serializeData, nil
// }

// func (embeddedTransferTransactionV1 EmbeddedTransferTransactionV1) Serialize() ([]byte, error) {
// 	// serialize inner common tx attrs
// 	serializeData, err := embeddedTransferTransactionV1.EmbeddedTransaction.serialize()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// serialize attrs
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.recipient.Bytes()...)
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.messageLength.Bytes()...)
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.mosaicsCount.Byte())
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.transferTransactionBodyReserved_1)
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.transferTransactionBodyReserved_2[:]...)

// 	// serialize mosiacs
// 	for _, mosaic := range embeddedTransferTransactionV1.mosaics {
// 		serializeData = append(serializeData, mosaic.Bytes()...)
// 	}

// 	// serialize message
// 	serializeData = append(serializeData, embeddedTransferTransactionV1.message.Bytes()...)

// 	return serializeData, nil
// }
