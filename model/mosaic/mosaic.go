package mosaic

import (
	"github.com/karriz-dev/symbol-sdk/model/decimal"
)

type Mosaic struct {
	MosaicId decimal.UInt64 // mosaic id
	Amount   decimal.UInt64 // amount of mosaic
}

type MosaicsCount decimal.UInt8

func (m Mosaic) Bytes() []byte {
	return append(m.MosaicId.Bytes(), m.Amount.Bytes()...)
}
