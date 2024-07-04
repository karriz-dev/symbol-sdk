package symbolsdk

import (
	"testing"

	"github.com/karriz-dev/symbol-sdk/common"
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
