package tx

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/network"
	"golang.org/x/crypto/sha3"
)

type AggregateBondedTransactionV2 struct {
	BaseTransaction

	aggregateTxMerkleRoot               Hash           // Hash of the aggregate's transaction.
	payloadSize                         decimal.UInt32 // Transaction payload size in bytes. This is the total number of bytes occupied by all embedded transactions, including any padding present.
	aggregateTransactionHeaderReserved1 decimal.UInt32 // reserved padding value
	transactions                        []Transaction  // Embedded transaction data. Transactions are variable-sized and the total payload size is in bytes. Embedded transactions cannot be aggregates.
}

func NewAggregateBondedTransactionV2(network network.Network, maxFee decimal.UInt64, deadline decimal.UInt64, signer account.PublicKey) AggregateBondedTransactionV2 {
	baseTx := BaseTransaction{
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

	baseTx.SetBaseSize(40, false)

	return AggregateBondedTransactionV2{
		BaseTransaction:                     baseTx,
		payloadSize:                         decimal.NewUInt32(0),
		aggregateTransactionHeaderReserved1: decimal.NewUInt32(0),
	}
}

func (tx *AggregateBondedTransactionV2) Transactions(transactions []Transaction) *AggregateBondedTransactionV2 {
	if len(transactions) <= 0 {
		return tx
	}

	tx.payloadSize = decimal.NewUInt32(0)

	for _, innerTx := range transactions {
		tx.payloadSize.Add(innerTx.Size())
	}

	hash, err := MerkleRootHash(transactions)
	if err != nil {
		return tx
	}

	tx.size.Add(tx.payloadSize.Value())
	tx.aggregateTxMerkleRoot = hash
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
	serializeData = append(serializeData, tx.aggregateTxMerkleRoot[:]...)
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
	return tx.aggregateTxMerkleRoot
}

func (tx AggregateBondedTransactionV2) Hash(generationHashSeed []byte) Hash {
	hasher := sha3.New256()

	// TODO :: need error check
	payload, _ := tx.Payload()

	hasher.Write(tx.signature[:])    // signature
	hasher.Write(tx.signer[:])       // signer public key
	hasher.Write(generationHashSeed) // generationhashssed
	hasher.Write(payload)            // tx payload

	hashedBytes := hasher.Sum(nil)

	return Hash(hashedBytes)
}

func (tx AggregateBondedTransactionV2) Payload() (Payload, error) {
	serializedBytes, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	return Payload(serializedBytes[TransactionHeaderSize:(TransactionHeaderSize + AggregateHashedSize)]), nil
}
