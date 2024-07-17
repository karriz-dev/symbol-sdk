package factory_test

import (
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/factory"
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/tx"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

var aliceAccount account.Account

func init() {
	// for testing fixed account
	account, err := account.NewRandomAccount(network.TESTNET)
	if err != nil {
		panic(err)
	}

	aliceAccount = account
}

func TestTransferTransactionV1(t *testing.T) {
	// set transaction factory
	transactionFactory := factory.NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(aliceAccount.PublicKey).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1(false)

	transferTx.
		Recipient(aliceAccount.Address)

	// sign
	sign, err := transactionFactory.Sign(transferTx, aliceAccount.PrivateKey)
	require.NoError(t, err)

	// attach sign
	transferTx.AttachSignature(sign)

	t.Logf("signature: %s", sign.Hex())
}

func TestEmbeddedTransaction(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1 (isEmbedded = true)
	embeddedTransferTx := txFactory.
		Signer(aliceAccount.PublicKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(aliceAccount.Address)

	innerTxList := []tx.Transaction{embeddedTransferTx, embeddedTransferTx, embeddedTransferTx}

	t.Logf("inner tx count: %d", len(innerTxList))

	for i, innerTx := range innerTxList {
		innerTxSerializedBytes, err := innerTx.Serialize()
		require.NoError(t, err)

		t.Logf("No.%d innerTxSerializedBytes: %+v", i, innerTxSerializedBytes)
	}
}

func TestAggregateTransaction(t *testing.T) {
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	embeddedTransferTx := txFactory.
		Signer(aliceAccount.PublicKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(aliceAccount.Address)

	innerTxList := []tx.Transaction{embeddedTransferTx, embeddedTransferTx, embeddedTransferTx}

	aggregateTx := txFactory.
		Signer(aliceAccount.PublicKey).
		AggregateBondedTransactionV2()

	aggregateTx.Transactions(innerTxList)

	t.Logf("aggregateTx: %+v", aggregateTx)
}
