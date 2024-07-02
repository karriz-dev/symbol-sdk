package symbolsdk

import (
	"symbol-sdk/tx"
	"symbol-sdk/types"
)

type Utils struct {
	F func()
}

type SymbolFacade struct {
	TransactionFactory tx.TransactionFactory

	MainNet types.NetworkType
	TestNet types.NetworkType
}

// Supported symbol network type
//
//	mainnet: symbol main network (104)
//	testnet: symbol sai test network (152)
func NewSymbolFacade(networkType string) *SymbolFacade {
	var network types.NetworkType
	switch networkType {
	case "mainnet":
		network = types.MAINNET
	case "testnet":
		network = types.TESTNET
	default:
		network = types.TESTNET
	}
	return &SymbolFacade{
		TransactionFactory: *tx.NewTransactionFactory(network),
		MainNet:            types.MAINNET,
		TestNet:            types.TESTNET,
	}
}
