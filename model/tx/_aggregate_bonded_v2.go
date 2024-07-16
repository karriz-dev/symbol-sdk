package tx

import (
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/merkle"
)

type AggregateBondedTransactionV2 struct {
	Transaction

	transactionHash                     Hash           // Hash of the aggregate's transaction.
	payloadSize                         decimal.UInt32 // Transaction payload size in bytes. This is the total number of bytes occupied by all embedded transactions, including any padding present.
	aggregateTransactionHeaderReserved1 int            // reserved padding value
	transactions                        []Transaction  // Embedded transaction data. Transactions are variable-sized and the total payload size is in bytes. Embedded transactions cannot be aggregates.
}

// func (transactionFactory *factory.TransactionFactory) AggregateBondedTransactionV2() AggregateBondedTransactionV2 {
// 	tx := BaseTransaction{
// 		size:                            40,
// 		version:                         0x02,
// 		network:                         transactionFactory.network,
// 		txType:                          0x4241,
// 		fee:                             transactionFactory.maxFee,
// 		deadline:                        transactionFactory.deadline,
// 		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
// 		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
// 		signer:                          transactionFactory.signer,
// 		isEmbedded:                      false,
// 	}

// 	return AggregateBondedTransactionV2{
// 		Transaction:                         tx,
// 		aggregateTransactionHeaderReserved1: 0,
// 	}
// }

func (tx *AggregateBondedTransactionV2) Transactions(transactions []Transaction) *AggregateBondedTransactionV2 {
	if len(transactions) <= 0 {
		return tx
	}

	for _, innerTx := range transactions {
		tx.payloadSize.Add(innerTx.Size())
	}

	hash, err := merkle.MerkleRootHash(transactions)
	if err != nil {
		return tx
	}

	tx.transactionHash = hash
	tx.transactions = transactions

	return tx
}

func (tx AggregateBondedTransactionV2) Serialize() []byte {
	// serialize inner common tx attrs
	serializeData := tx.Transaction.Serialize()
	if len(serializeData) <= 0 {
		return nil
	}

	// serialize attrs
	serializeData = append(serializeData, common.Bytes(tx.transactionHash)...)
	serializeData = append(serializeData, common.Bytes(tx.payloadSize)...)
	serializeData = append(serializeData, common.Bytes(tx.aggregateTransactionHeaderReserved1)...)

	for _, innerTx := range tx.transactions {
		serializeData = append(serializeData, innerTx.Serialize()...)
	}

	return serializeData
}

func (tx AggregateBondedTransactionV2) MerkleRootHash() common.Hash {
	return tx.transactionHash
}

func (tx AggregateBondedTransactionV2) Hash() common.Hash {
	return common.Hash{}
}
