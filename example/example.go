package example

import (
	"errors"
	"os"
	"strings"

	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/network"
)

const (
	SymbolTestNetworkUrl = "https://001-sai-dual.symboltest.net:3001"
)

var AliceAccount account.Account
var BobAccount account.Account
var ServiceProviderAccount account.Account

func init() {
	data, err := os.ReadFile("./private_keys.txt")
	if err != nil {
		panic(err)
	}

	privateKeys := strings.Split(string(data), "\n")
	accountList := []account.Account{}

	for _, privateKeyHexString := range privateKeys {
		privateKey, err := account.PrivateKeyFromHex(privateKeyHexString)
		if err != nil {
			panic(err)
		}

		acc, err := account.AccountFromPrivateKey(privateKey, network.TESTNET)
		if err != nil {
			panic(err)
		}

		accountList = append(accountList, acc)
	}

	if len(accountList) != 3 {
		panic(errors.New("private keys must be 3"))
	}

	AliceAccount = accountList[0]
	BobAccount = accountList[1]
	ServiceProviderAccount = accountList[2]
}
