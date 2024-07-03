package tx

import (
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

func TestTransactionFactory(t *testing.T) {
	// set transaction factory
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1()

	t.Logf("transferTx: %+v", transferTx)
}

func TestTransactionSerialize(t *testing.T) {
	// set transaction factory
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// alice keyPair
	aliceKeyPair, err := common.HexToKeyPair("38FB967C5427C6D4CAF9BEFBF8B80B0D139BC55374643F76A13325F852C09DFD")
	require.NoError(t, err)

	// bob keyPair
	bobKeyPair, err := common.HexToKeyPair("B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EA72CF28AAD1")
	require.NoError(t, err)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(aliceKeyPair).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1()

	// make transfer_transaction_v1 [alice -> bob xym 1]
	serializedData, err := transferTx.
		Recipient(common.PublicKeyToAddress(bobKeyPair.PublicKey, network.TESTNET)).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).Serialize()

	require.NoError(t, err)

	txHex := common.BytesToHex(serializedData)
	t.Logf("tx hex: %s", txHex)

	txHash, err := common.BytesToHash(serializedData)
	require.NoError(t, err)

	t.Logf("tx hash: %s", common.BytesToHex(txHash[:]))
}

func TestTransactionSign(t *testing.T) {
	// set transaction factory
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// alice keyPair
	aliceKeyPair, err := common.HexToKeyPair("38FB967C5427C6D4CAF9BEFBF8B80B0D139BC55374643F76A13325F852C09DFD")
	require.NoError(t, err)

	// bob keyPair
	bobKeyPair, err := common.HexToKeyPair("B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EA72CF28AAD1")
	require.NoError(t, err)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(aliceKeyPair).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1()

	err = transferTx.
		Recipient(common.PublicKeyToAddress(bobKeyPair.PublicKey, network.TESTNET)).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).Sign()
	require.NoError(t, err)

	serializeData, err := transferTx.Serialize()
	require.NoError(t, err)

	t.Logf("tx hex: %s", common.BytesToHex(serializeData))
}
