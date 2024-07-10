package tx

import (
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

func TestCommonTransaction(t *testing.T) {
	// set transaction factory
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(common.HexToPublicKey("")).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1(false)

	// sign
	sign, err := transactionFactory.Sign(transferTx, common.HexToPrivateKey(""))
	require.NoError(t, err)

	// attach sign
	transferTx.AttachSignature(sign)

	// logging
	t.Logf("signature hex: %s", common.BytesToHex(sign[:]))
	t.Logf("transferTx: %+v", transferTx)
}

func TestEmbeddedTransaction(t *testing.T) {
	// set transaction factory
	txFactory := NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1 (isEmbedded = true)
	innerTxList := make([]ITransaction, 0)

	embeddedTransferTx := txFactory.
		Signer(common.HexToPublicKey("14BCAB6B7D2358F9C31A40969D0ACC48125831C0390135CA35AEA282D2A54AAC")).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI")).
		Message("Hello, I'm Embedded Tx")

	innerTxList = append(innerTxList, embeddedTransferTx)
	innerTxList = append(innerTxList, embeddedTransferTx)
	innerTxList = append(innerTxList, embeddedTransferTx)

	t.Logf("inner tx count: %d", len(innerTxList))
	t.Logf("inner tx list: %+v", innerTxList)

	for _, etx := range innerTxList {
		serializeBytes, _ := etx.Serialize()

		t.Logf("serializeBytes: %+v", serializeBytes)
	}
}
