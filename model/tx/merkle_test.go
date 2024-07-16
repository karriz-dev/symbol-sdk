package tx_test

import (
	"testing"

	"github.com/karriz-dev/symbol-sdk/factory"
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/tx"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

var testAccount account.Account

func init() {
	account, err := account.NewRandomAccount(network.TESTNET)
	if err != nil {
		panic(err)
	}

	testAccount = account
}

func TestCalcMerkleHashOneTransaction(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1
	embeddedTransferTx := txFactory.
		Signer(testAccount.PublicKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(testAccount.Address)

	merkleRootHash, err := tx.MerkleRootHash([]tx.Transaction{embeddedTransferTx})
	require.NoError(t, err)

	t.Log(merkleRootHash)
}

func TestCalcMerkleHashEvenTransactions(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1
	embeddedTransferTx := txFactory.
		Signer(testAccount.PublicKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(testAccount.Address)

	merkleRootHash, err := tx.MerkleRootHash([]tx.Transaction{embeddedTransferTx, embeddedTransferTx})
	require.NoError(t, err)

	t.Log(merkleRootHash)
}

func TestCalcMerkleHashOddTransactions(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1
	embeddedTransferTx := txFactory.
		Signer(testAccount.PublicKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(testAccount.Address)

	merkleRootHash, err := tx.MerkleRootHash([]tx.Transaction{embeddedTransferTx, embeddedTransferTx, embeddedTransferTx})
	require.NoError(t, err)

	t.Log(merkleRootHash)
}
