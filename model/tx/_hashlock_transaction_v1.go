package tx

import (
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/model/mosaic"
	"github.com/karriz-dev/symbol-sdk/types"
)

type HashLockTransactionV1 struct {
	Transaction

	mosaic   mosaic.Mosaic       // Locked mosaic
	duration types.BlockDuration // Number of blocks for which a lock should be valid. The default maximum is 48h (See the maxHashLockDuration network property).
	hash     Hash                // Hash of the AggregateBondedTransaction to be confirmed before unlocking the mosaics.
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

func (tx *HashLockTransactionV1) Mosaic(mosaic common.Mosaic) *HashLockTransactionV1 {
	tx.mosaic = mosaic

	return tx
}

func (tx *HashLockTransactionV1) LockDuration(duration types.BlockDuration) *HashLockTransactionV1 {
	tx.duration = duration

	return tx
}

func (tx *HashLockTransactionV1) ParentHash(hash common.Hash) *HashLockTransactionV1 {
	tx.hash = hash

	return tx
}

func (tx HashLockTransactionV1) Serialize() []byte {
	// serialize inner common tx attrs
	serializeData := tx.Transaction.Serialize()
	if len(serializeData) <= 0 {
		return nil
	}

	// serialize attrs
	serializeData = append(serializeData, common.Bytes(tx.mosaic)...)
	serializeData = append(serializeData, common.Bytes(tx.duration)...)
	serializeData = append(serializeData, common.Bytes(tx.hash)...)

	return serializeData
}
