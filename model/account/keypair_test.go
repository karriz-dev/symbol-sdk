package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrivateKeyFromHex(t *testing.T) {
	privateKey, err := PrivateKeyFromHex("B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EA72CF28AAD2")
	require.NoError(t, err)

	assert.Equal(t, "B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EA72CF28AAD2", privateKey.Hex())
}

func TestPrivateKeyFromHexWithError(t *testing.T) {
	privateKey, err := PrivateKeyFromHex("B4EEF6D6A27004B5D9AF38B03234802E27B4CEFEC63CD8AD2B33EACF28AAD1")

	assert.Error(t, err)
	assert.Empty(t, privateKey)
}
