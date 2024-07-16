package common

type Hash [32]byte
type Signature [64]byte
type CoSignature struct {
	Version   uint16
	Signer    PublicKey
	Signature Signature
}
