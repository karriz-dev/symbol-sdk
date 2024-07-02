package common

import (
	"encoding/binary"
)

type Mosaic struct {
	MosaicId uint64 // mosaic id
	Amount   uint64 // amount of mosaic
}

type MosaicsCount uint8

func (mosaic Mosaic) Bytes() []byte {
	mosaicBytes := make([]byte, 8)
	amountBytes := make([]byte, 8)

	binary.LittleEndian.PutUint64(mosaicBytes, mosaic.MosaicId)
	binary.LittleEndian.PutUint64(amountBytes, mosaic.Amount)

	return append(mosaicBytes, amountBytes...)
}

func (mosaicsCount MosaicsCount) Byte() byte {
	return byte(mosaicsCount)
}
