package signature

import (
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
)

type CoSignature struct {
	version   decimal.UInt16
	signer    account.PublicKey
	signature Signature
}

func NewCoSignature(version decimal.UInt16, signer account.PublicKey, signature Signature) CoSignature {
	return CoSignature{
		version:   version,
		signer:    signer,
		signature: signature,
	}
}
