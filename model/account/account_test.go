package account

import (
	"testing"

	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	account, err := NewRandomAccount(network.TESTNET)
	require.NoError(t, err)

	t.Logf("%+v", account)
}

func TestAccountFromPrivateKey(t *testing.T) {
	// privateKey hex for testing
	alicePrivateKey, err := PrivateKeyFromHex("AE91E63911B08A5E53D9EAB9FA52D76E2A2FC95D2BB965820F3AF03C26185BFD")
	require.NoError(t, err)

	account, err := AccountFromPrivateKey(alicePrivateKey, network.TESTNET)
	require.NoError(t, err)

	assert.Equal(t, account.PrivateKey.Hex(), "AE91E63911B08A5E53D9EAB9FA52D76E2A2FC95D2BB965820F3AF03C26185BFD")
	assert.Equal(t, account.PublicKey.Hex(), "8E734BAF8D595F27AAA5DAD79999536A0B00573C6503773199F9173B8B23C293")
	assert.Equal(t, account.EncodedAddress(), "TAUWSCUU7ZBEZMRU66ATIAMDJE55IC56CCBDPIQ")
}
