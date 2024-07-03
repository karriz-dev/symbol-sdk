package symbolsdk

import (
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/karriz-dev/symbol-sdk/tx"
)

type SymbolFacade struct {
	TransactionFactory tx.TransactionFactory
	Network            network.Network
}

// Supported symbol network type
//
//	mainnet: symbol main network (104)
//	testnet: symbol sai test network (152)
func NewSymbolFacade(networkType string) *SymbolFacade {
	var symbolNetwork network.Network
	switch networkType {
	case "mainnet":
		symbolNetwork = network.MAINNET
	case "testnet":
		symbolNetwork = network.TESTNET
	default:
		symbolNetwork = network.TESTNET
	}
	return &SymbolFacade{
		TransactionFactory: *tx.NewTransactionFactory(symbolNetwork),
		Network:            symbolNetwork,
	}
}
