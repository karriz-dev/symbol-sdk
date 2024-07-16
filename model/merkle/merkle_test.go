package merkle

import (
	"testing"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/factory"
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/tx"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcMerkleHashOneTransaction(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// testing account
	testAccountPubKey, err := account.PublicKeyFromHex("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")
	require.NoError(t, err)

	// create embedded TransferTransactionV1
	embeddedTransferTx := txFactory.
		Signer(testAccountPubKey).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

	merkleRootHash, err := MerkleRootHash([]tx.Transaction{embeddedTransferTx})
	require.NoError(t, err)

	assert.Equal(t, "B701124087B58ACC62196EC81B8437E0D5D1A064B47BD719E0ECA74452031894", common.BytesToHex(merkleRootHash[:]))
}

func TestCalcMerkleHashEvenTransactions(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1, 2
	transactions := make([]tx.Transaction, 0)
	embeddedTransferTx := txFactory.
		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
		TransferTransactionV1(true)

	embeddedTransferTx.
		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

	transactions = append(transactions, embeddedTransferTx)
	transactions = append(transactions, embeddedTransferTx)

	merkleRootHash, err := MerkleRootHash(transactions)
	require.NoError(t, err)

	assert.Equal(t, "ACDD2D30AFED90C0ADDE938A40BE547F76A932BF93A79C23351441B276857FBB", common.BytesToHex(merkleRootHash[:]))
}

func TestCalcMerkleHashOddTransactions(t *testing.T) {
	// set transaction factory
	txFactory := factory.NewTransactionFactory(network.TESTNET)

	// create embedded TransferTransactionV1, 2
	transactions := make([]tx.Transaction, 0)
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

	assert.Equal(t, "06FAD3281362B48293C077FB3760A13B86FADB25941B14EEC2A5E9AFE2048C07", common.BytesToHex(merkleRootHash[:]))
}
