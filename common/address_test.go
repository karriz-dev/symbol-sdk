package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexToAddress(t *testing.T) {
	address, err := HexToAddress("98ECD29F49074A6DAFAB99B0897661E7FB323735C2782809")
	require.NoError(t, err)

	t.Log(address.Encode())
}
