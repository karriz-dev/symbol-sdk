package symbolsdk

import (
	"symbol-sdk/common"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var aliceKeyPair common.KeyPair
var bobKeyPair common.KeyPair

func init() {
	aliceKeyPair, _ = common.NewKeyPair()
	bobKeyPair, _ = common.NewKeyPair()
}

func TestNewSymbolFacade(t *testing.T) {
	symbolFacade := NewSymbolFacade("testnet")

	t.Log(symbolFacade)
}

func TestTransactionFactory(t *testing.T) {
	symbolFacade := NewSymbolFacade("testnet")

	transferTx := symbolFacade.TransactionFactory.
		Signer(aliceKeyPair).
		MaxFee(1_000000).
		Deadline(time.Minute * 10).
		TransferTransactionV1()

	t.Log(transferTx)
}

func TestTransactionSign(t *testing.T) {
	symbolFacade := NewSymbolFacade("testnet")

	transferTx := symbolFacade.TransactionFactory.
		Signer(aliceKeyPair).
		MaxFee(1_000000).
		Deadline(time.Minute * 10).
		TransferTransactionV1()

	err := transferTx.
		Recipient(common.PublicKeyToAddress(bobKeyPair.PublicKey, symbolFacade.TestNet)).
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
