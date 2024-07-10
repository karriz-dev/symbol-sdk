package tx

import (
	"testing"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

func TestCalcMerkleHashOneTransaction(t *testing.T) {
	// set transaction factory
	txFactory := NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1
	embeddedTransferTx := txFactory.
		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

	merkleRootHash, err := MerkleRootHash([]ITransaction{embeddedTransferTx})
	require.NoError(t, err)

	t.Log(common.BytesToHex(merkleRootHash[:]))
}

func TestCalcMerkleHashEvenTransactions(t *testing.T) {
	// set transaction factory
	txFactory := NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1, 2
	transactions := make([]ITransaction, 0)
	embeddedTransferTx := txFactory.
		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

	transactions = append(transactions, embeddedTransferTx)
	transactions = append(transactions, embeddedTransferTx)

	merkleRootHash, err := MerkleRootHash(transactions)
	require.NoError(t, err)

	t.Log(common.BytesToHex(merkleRootHash[:]))
}

func TestCalcMerkleHashOddTransactions(t *testing.T) {
	// set transaction factory
	txFactory := NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1, 2
	transactions := make([]ITransaction, 0)
	embeddedTransferTx := txFactory.
		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

	transactions = append(transactions, embeddedTransferTx)
	transactions = append(transactions, embeddedTransferTx)
	transactions = append(transactions, embeddedTransferTx)

	merkleRootHash, err := MerkleRootHash(transactions)
	require.NoError(t, err)

	t.Log(common.BytesToHex(merkleRootHash[:]))
}
