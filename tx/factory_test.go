package tx

import (
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
)

func TestTransactionFactory(t *testing.T) {
	// set transaction factory
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(common.HexToPublicKey("14BCAB6B7D2358F9C31A40969D0ACC48125831C0390135CA35AEA282D2A54AAC")).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1()

	t.Logf("transferTx: %+v", transferTx)
}

func TestEmbeddedTransactionFactory(t *testing.T) {
	// set transaction factory
	embeddedTransactionFactory := NewEmbeddedTransactionFactory(network.TESTNET)

	// create EmbeddedTransferTransactionV1
	embeddedTransactionList := make([]IEmbeddedTransaction, 0)

	embeddedTransferTx := embeddedTransactionFactory.
		Signer(common.HexToPublicKey("14BCAB6B7D2358F9C31A40969D0ACC48125831C0390135CA35AEA282D2A54AAC")).
		EmbeddedTransferTransactionV1()

	aliceAddress, _ := common.DecodeAddress("TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI")

	embeddedTransferTx.
		Recipient(aliceAddress).
		Message("Hello, I'm Embedded Tx")

	embeddedTransactionList = append(embeddedTransactionList, embeddedTransferTx)
	embeddedTransactionList = append(embeddedTransactionList, embeddedTransferTx)
	embeddedTransactionList = append(embeddedTransactionList, embeddedTransferTx)

	t.Logf("embeddedTransactionList Count: %d", len(embeddedTransactionList))
	t.Logf("embeddedTransactionList: %+v", embeddedTransactionList)

	for _, etx := range embeddedTransactionList {
		serializeBytes, _ := etx.Serialize()

		t.Logf("serializeBytes: %+v", serializeBytes)
	}
}
