package tx

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/network"
)

type AggregateBondedTransactionV2 struct {
	BaseTransaction

	transactionHash                     Hash           // Hash of the aggregate's transaction.
	payloadSize                         decimal.UInt32 // Transaction payload size in bytes. This is the total number of bytes occupied by all embedded transactions, including any padding present.
	aggregateTransactionHeaderReserved1 decimal.UInt32 // reserved padding value
	transactions                        []Transaction  // Embedded transaction data. Transactions are variable-sized and the total payload size is in bytes. Embedded transactions cannot be aggregates.
}

func NewAggregateBondedTransactionV2(network network.Network, maxFee decimal.UInt64, deadline decimal.UInt64, signer account.PublicKey) AggregateBondedTransactionV2 {
	tx := BaseTransaction{
		size:                            decimal.NewUInt32(40),
		version:                         decimal.NewUInt8(0x02),
		network:                         network,
		txType:                          decimal.NewUInt16(0x4241),
		fee:                             maxFee,
		deadline:                        deadline,
		verifiableEntityHeaderReserved1: decimal.NewUInt32(0),
		entityBodyReserved1:             decimal.NewUInt32(0),
		signer:                          signer,
		isEmbedded:                      false,
	}
	return AggregateBondedTransactionV2{
		BaseTransaction:                     tx,
		payloadSize:                         decimal.NewUInt32(0),
		aggregateTransactionHeaderReserved1: decimal.NewUInt32(0),
	}
}

func (tx *AggregateBondedTransactionV2) Transactions(transactions []Transaction) *AggregateBondedTransactionV2 {
	if len(transactions) <= 0 {
		return tx
	}

	for _, innerTx := range transactions {
		tx.payloadSize.Add(innerTx.Size())
	}

	hash, err := MerkleRootHash(transactions)
	if err != nil {
		return tx
	}

	tx.transactionHash = hash
	tx.transactions = transactions

	return tx
}

func (tx AggregateBondedTransactionV2) Serialize() ([]byte, error) {
	// serialize inner common tx attrs
	serializeData, err := tx.BaseTransaction.Serialize()
	if err != nil {
		return nil, err
	}

	// serialize attrs
	serializeData = append(serializeData, tx.transactionHash[:]...)
	serializeData = append(serializeData, tx.payloadSize.Bytes()...)
	serializeData = append(serializeData, tx.aggregateTransactionHeaderReserved1.Bytes()...)

	for _, innerTx := range tx.transactions {
		innerTxSerializedBytes, err := innerTx.Serialize()
		if err != nil {
			return nil, err
		}

		serializeData = append(serializeData, innerTxSerializedBytes...)
	}

	return serializeData, nil
}

func (tx AggregateBondedTransactionV2) MerkleRootHash() Hash {
	return tx.transactionHash
}

func (tx AggregateBondedTransactionV2) Hash() Hash {
	return Hash{}
}
