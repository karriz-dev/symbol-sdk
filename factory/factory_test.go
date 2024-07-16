package factory

import (
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/model/account"
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
	transactionFactory := NewTransactionFactory(network.TESTNET)

	// create TransferTransactionV1
	transferTx := transactionFactory.
		Signer(aliceAccount.PublicKey).
		MaxFee(1_000000).
		Deadline(time.Hour * 2).
		TransferTransactionV1(false)

	// sign
	sign, err := transactionFactory.Sign(transferTx, aliceAccount.PrivateKey)
	require.NoError(t, err)

	// attach sign
	transferTx.AttachSignature(sign)

	// logging
	t.Logf("signature hex: %s", sign.Hex())
	t.Logf("transferTx: %+v", transferTx)
}

// func TestEmbeddedTransaction(t *testing.T) {
// 	// set transaction factory
// 	txFactory := NewTransactionFactory(network.TESTNET)

// 	// create TransferTransactionV1 (isEmbedded = true)
// 	innerTxList := make([]ITransaction, 0)

// 	embeddedTransferTx := txFactory.
// 		Signer(common.HexToPublicKey("14BCAB6B7D2358F9C31A40969D0ACC48125831C0390135CA35AEA282D2A54AAC")).
// 		TransferTransactionV1(true)

// 	embeddedTransferTx.
// 		Recipient(common.DecodeAddress("TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI")).
// 		Message("Hello, I'm Embedded Tx")

// 	innerTxList = append(innerTxList, embeddedTransferTx)
// 	innerTxList = append(innerTxList, embeddedTransferTx)
// 	innerTxList = append(innerTxList, embeddedTransferTx)

// 	t.Logf("inner tx count: %d", len(innerTxList))
// 	t.Logf("inner tx list: %+v", innerTxList)

// 	for _, etx := range innerTxList {
// 		t.Logf("serializeBytes: %+v", etx.Serialize())
// 	}
// }

// func TestAggregateTransaction(t *testing.T) {
// 	txFactory := NewTransactionFactory(network.TESTNET)

// 	innerTxList := make([]ITransaction, 0)

// 	embeddedTransferTx := txFactory.
// 		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
// 		TransferTransactionV1(true)

// 	embeddedTransferTx.
// 		Recipient(common.DecodeAddress("TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA"))

// 	innerTxList = append(innerTxList, embeddedTransferTx)
// 	innerTxList = append(innerTxList, embeddedTransferTx)
// 	innerTxList = append(innerTxList, embeddedTransferTx)

// 	aggregateTx := txFactory.
// 		Signer(common.HexToPublicKey("8B74E74D1EC9496E5CDCE56BE21330668CA6EAAFE73CC033E7C35D6B4DEE2179")).
// 		AggregateBondedTransactionV2()

// 	aggregateTx.Transactions(innerTxList)

// 	t.Logf("%+v", aggregateTx)
// }

// func TestSwapWithServiceProvider(t *testing.T) {
// 	// if you need test this function, plz input 3 account's private key
// 	alice := common.HexToPrivateKey("B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EA72CF28AAD1")
// 	bob := common.HexToPrivateKey("CBEE65061055E60528264EAA0496BDE4DDC74E83F529C818DFA18E9BF253E0E4")
// 	serviceProvider := common.HexToPrivateKey("38FB967C5427C6D4CAF9BEFBF8B80B0D139BC55374643F76A13325F852C09DFD")

// 	aliceAddress := common.PublicKeyToAddress(alice.PublicKey(), network.TESTNET)
// 	bobAddress := common.PublicKeyToAddress(bob.PublicKey(), network.TESTNET)
// 	serivceProviderAddress := common.PublicKeyToAddress(serviceProvider.PublicKey(), network.TESTNET)

// 	// require 3 account's private key & address is valid
// 	require.NotEmpty(t, alice)
// 	require.NotEmpty(t, bob)
// 	require.NotEmpty(t, serviceProvider)
// 	require.NotEmpty(t, aliceAddress)
// 	require.NotEmpty(t, bobAddress)
// 	require.NotEmpty(t, serivceProviderAddress)

// 	// new transaction factory
// 	txFactory := NewTransactionFactory(network.TESTNET)

// 	// check tx factory isn't nil
// 	require.NotNil(t, txFactory)

// 	// 1. create 3 inner tx (self tx serviceProvider, alice->bob 1xym, bob->alice 1xym)
// 	tx1 := txFactory.
// 		Signer(serviceProvider.PublicKey()).
// 		TransferTransactionV1(true)
// 	tx1.Recipient(serivceProviderAddress)

// 	tx2 := txFactory.
// 		Signer(alice.PublicKey()).
// 		TransferTransactionV1(true)
// 	tx2.Recipient(bobAddress).
// 		Mosaics([]common.Mosaic{
// 			{
// 				MosaicId: 0x72C0212E67A08BCE,
// 				Amount:   1_000000,
// 			},
// 		})

// 	tx3 := txFactory.
// 		Signer(bob.PublicKey()).
// 		TransferTransactionV1(true)
// 	tx3.Recipient(aliceAddress).
// 		Mosaics([]common.Mosaic{
// 			{
// 				MosaicId: 0x72C0212E67A08BCE,
// 				Amount:   1_000000,
// 			},
// 		})

// 	innerTxList := []ITransaction{tx1, tx2, tx3}

// 	t.Log("innerTxList:", innerTxList)

// 	// 2. create aggregate_bonded tx with 3 inner txs & sign (Signer: serviceProvider)
// 	aggregateTx := txFactory.
// 		Signer(serviceProvider.PublicKey()).
// 		AggregateBondedTransactionV2()

// 	aggregateTx.Transactions(innerTxList)

// 	aggregateTxSign, err := txFactory.Sign(aggregateTx, serviceProvider)
// 	require.NoError(t, err)

// 	aggregateTx.AttachSignature(aggregateTxSign)

// 	t.Log("aggregateTx.MerkleRootHash(Hex):", common.BytesToHex(common.Bytes(aggregateTx.MerkleRootHash())))

// 	// 3. create hashlock tx with aggregate_bonded & sign (Signer: serviceProvider)
// 	hashlockTx := txFactory.
// 		Signer(serviceProvider.PublicKey()).
// 		HashLockTransactionV1(false)

// 	hashlockTx.
// 		Mosaic(common.Mosaic{
// 			MosaicId: 0x72C0212E67A08BCE,
// 			Amount:   10_000000,
// 		}).
// 		LockDuration(types.BlockDuration(10)).
// 		ParentHash(aggregateTx.MerkleRootHash())

// 	hashlockTxSign, err := txFactory.Sign(hashlockTx, serviceProvider)
// 	require.NoError(t, err)

// 	hashlockTx.AttachSignature(hashlockTxSign)

// 	t.Log("hashlockTx.Serialize:", hashlockTx.Serialize())
// }
