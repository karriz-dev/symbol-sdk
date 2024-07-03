package tx

import "github.com/karriz-dev/symbol-sdk/common"

type MosaicDefinitionV1 struct {
	Transaction

	recipient common.Address
	mosaics   []common.Mosaic
	message   common.Message

	messageLength common.MessageLength
	mosaicsCount  common.MosaicsCount

	transferTransactionBodyReserved_1 byte
	transferTransactionBodyReserved_2 [4]byte
}
