package types

import (
	"encoding/binary"
)

type TransactionSize uint32 // transaction size
type TransactionType uint16 // type of transaction
type MaxFee uint64          // maximum fee amount
type Deadline uint64        // transaction sign deadline

func (transactionSize TransactionSize) Bytes() []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(transactionSize))

	return result
}

func (transactionType TransactionType) Bytes() []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, uint16(transactionType))

	return result
}

func (maxFee MaxFee) Bytes() []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint32(result, uint32(maxFee))

	return result
}

func (deadline Deadline) Bytes() []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(deadline))

	return result
}
