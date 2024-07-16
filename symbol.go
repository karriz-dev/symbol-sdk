package symbolsdk

import (
	"github.com/karriz-dev/symbol-sdk/factory"
	"github.com/karriz-dev/symbol-sdk/network"
)

type SymbolFacade struct {
	TransactionFactory factory.TransactionFactory
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
		TransactionFactory: *factory.NewTransactionFactory(symbolNetwork),
		Network:            symbolNetwork,
	}
}
