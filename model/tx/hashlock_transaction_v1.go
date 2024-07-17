package tx

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/mosaic"
	"github.com/karriz-dev/symbol-sdk/network"
)

type HashLockTransactionV1 struct {
	BaseTransaction

	mosaic   mosaic.Mosaic  // Locked mosaic
	duration decimal.UInt64 // Number of blocks for which a lock should be valid. The default maximum is 48h (See the maxHashLockDuration network property).
	hash     Hash           // Hash of the AggregateBondedTransaction to be confirmed before unlocking the mosaics.
}

func NewHashLockTransactionV1(network network.Network, maxFee decimal.UInt64, deadline decimal.UInt64, signer account.PublicKey, isEmbedded bool) HashLockTransactionV1 {
	baseTx := BaseTransaction{
		version:                         decimal.NewUInt8(0x01),
		network:                         network,
		txType:                          decimal.NewUInt16(0x4148),
		fee:                             maxFee,
		deadline:                        deadline,
		verifiableEntityHeaderReserved1: decimal.NewUInt32(0),
		entityBodyReserved1:             decimal.NewUInt32(0),
		signer:                          signer,
		isEmbedded:                      isEmbedded,
	}

	baseTx.SetBaseSize(56, isEmbedded)

	return HashLockTransactionV1{
		BaseTransaction: baseTx,
	}
}

func (tx *HashLockTransactionV1) Mosaic(mosaic mosaic.Mosaic) *HashLockTransactionV1 {
	tx.mosaic = mosaic

	return tx
}

func (tx *HashLockTransactionV1) LockDuration(duration decimal.UInt64) *HashLockTransactionV1 {
	tx.duration = duration

	return tx
}

func (tx *HashLockTransactionV1) ParentHash(hash Hash) *HashLockTransactionV1 {
	tx.hash = hash

	return tx
}

func (tx HashLockTransactionV1) Serialize() ([]byte, error) {
	// serialize inner common tx attrs
	serializeData, err := tx.BaseTransaction.Serialize()
	if err != nil {
		return nil, err
	}

	// serialize attrs
	serializeData = append(serializeData, tx.mosaic.Bytes()...)
	serializeData = append(serializeData, tx.duration.Bytes()...)
	serializeData = append(serializeData, tx.hash[:]...)

	return serializeData, nil
}

func (tx HashLockTransactionV1) Payload() (Payload, error) {
	serializedBytes, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	return Payload(serializedBytes[TransactionHeaderSize:]), nil
}
