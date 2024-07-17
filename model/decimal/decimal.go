package decimal

import (
	"encoding/binary"
	"strconv"

	"github.com/karriz-dev/symbol-sdk/model"
)

type UInt8 struct {
	model.Base
	value uint8
}

func (v UInt8) String() string {
	return strconv.FormatUint((uint64)(v.value), 10)
}

func (v UInt8) Bytes() []byte {
	return []byte{v.value}
}

func (v *UInt8) Add(addedValue uint8) {
	v.value += addedValue
}

func (v *UInt8) Value() uint8 {
	return v.value
}

type UInt16 struct {
	model.Base
	value uint16
}

func (v UInt16) String() string {
	return strconv.FormatUint((uint64)(v.value), 10)
}

func (v UInt16) Bytes() []byte {
	return binary.LittleEndian.AppendUint16(nil, uint16(v.value))
}

func (v *UInt16) Add(addedValue uint16) {
	v.value += addedValue
}

func (v *UInt16) Value() uint16 {
	return v.value
}

type UInt32 struct {
	model.Base
	value uint32
}

func (v UInt32) String() string {
	return strconv.FormatUint((uint64)(v.value), 10)
}

func (v UInt32) Bytes() []byte {
	return binary.LittleEndian.AppendUint32(nil, uint32(v.value))
}

func (v *UInt32) Add(addedValue uint32) {
	v.value += addedValue
}

func (v *UInt32) Value() uint32 {
	return v.value
}

type UInt64 struct {
	model.Base
	value uint64
}

func (v UInt64) String() string {
	return strconv.FormatUint((uint64)(v.value), 10)
}

func (v UInt64) Bytes() []byte {
	return binary.LittleEndian.AppendUint64(nil, uint64(v.value))
}

func (v *UInt64) Add(addedValue uint64) {
	v.value += addedValue
}

func (v *UInt64) Sub(addedValue uint64) {
	v.value -= addedValue
}

func (v *UInt64) Value() uint64 {
	return v.value
}

func NewUInt8(value uint8) UInt8 {
	return UInt8{
		value: value,
	}
}

func NewUInt16(value uint16) UInt16 {
	return UInt16{
		value: value,
	}
}

func NewUInt32(value uint32) UInt32 {
	return UInt32{
		value: value,
	}
}

func NewUInt64(value uint64) UInt64 {
	return UInt64{
		value: value,
	}
}
