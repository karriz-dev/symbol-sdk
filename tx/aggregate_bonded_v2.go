package tx

import "github.com/karriz-dev/symbol-sdk/common"

type AggregateBondedTransactionV2 struct {
	Transaction

	transactionHash                     common.Hash   // Hash of the aggregate's transaction.
	payloadSize                         int           // Transaction payload size in bytes. This is the total number of bytes occupied by all embedded transactions, including any padding present.
	aggregateTransactionHeaderReserved1 int           // reserved padding value
	transactions                        []Transaction // Embedded transaction data. Transactions are variable-sized and the total payload size is in bytes. Embedded transactions cannot be aggregates.
}

func (transactionFactory *TransactionFactory) AggregateBondedTransactionV2(isEmbedded bool) AggregateBondedTransactionV2 {
	tx := Transaction{
		size:                            40,
		version:                         0x02,
		network:                         transactionFactory.network,
		txType:                          0x4241,
		fee:                             transactionFactory.maxFee,
		deadline:                        transactionFactory.deadline,
		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
		signer:                          transactionFactory.signer,
		isEmbedded:                      isEmbedded,
	}

	return AggregateBondedTransactionV2{
		Transaction:                         tx,
		aggregateTransactionHeaderReserved1: 0,
	}                                                                  
}

func (tx *AggregateBondedTransactionV2) TransactionHash(transactionHash common.Hash) *AggregateBondedTransactionV2 {
	tx.transactionHash = transactionHash

	return tx
}

func (tx *AggregateBondedTransactionV2) PayloadSize(payloadSize int) *AggregateBondedTransactionV2 {
	tx.payloadSize = payloadSize

	return tx
}

func (tx *AggregateBondedTransactionV2) Transactions(transactions []Transaction) *AggregateBondedTransactionV2 {
	tx.transactions = transactions

	return tx
}
