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

func (transactionFactory *TransactionFactory) HashLockTransactionV1(isEmbedded bool) HashLockTransactionV1 {
	tx := Transaction{
		size:                            56,
		version:                         0x01,
		network:                         transactionFactory.network,
		txType:                          0x4148,
		fee:                             transactionFactory.maxFee,
		deadline:                        transactionFactory.deadline,
		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
		signer:                          transactionFactory.signer,
		isEmbedded:                      isEmbedded,
	}

	return HashLockTransactionV1{
		Transaction: tx,
	}
}

func (hashLockTransactionV1 *HashLockTransactionV1) Mosaic(mosaic common.Mosaic) *HashLockTransactionV1 {
	hashLockTransactionV1.mosaic = mosaic

	return hashLockTransactionV1
}

func (hashLockTransactionV1 *HashLockTransactionV1) LockDuration(duration time.Duration) *HashLockTransactionV1 {
	hashLockTransactionV1.duration = duration

	return hashLockTransactionV1
}

func (hashLockTransactionV1 *HashLockTransactionV1) Hash(hash common.Hash) *HashLockTransactionV1 {
	hashLockTransactionV1.hash = hash

	return hashLockTransactionV1
}

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
