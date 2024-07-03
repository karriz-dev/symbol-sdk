package network

import (
	"encoding/hex"
	"time"
)

var (
	MAINNET = Network{
		Name:               "mainnet",
		Type:               0x68,
		EpochTime:          time.Date(2021, time.March, 16, 0, 6, 25, 0, time.UTC),
		GenerationHashSeed: hexToBytes("57F7DA205008026C776CB6AED843393F04CD458E0AA2D9F1D5F31A402072B2D6"),
	}
	TESTNET = Network{
		Name:               "testnet",
		Type:               0x98,
		EpochTime:          time.Date(2022, time.October, 31, 21, 7, 47, 0, time.UTC),
		GenerationHashSeed: hexToBytes("49D6E1CE276A85B70EAFE52349AACCA389302E7A9754BCF1221E79494FC665A4"),
	}
)

type Network struct {
	Name               string    // Network name
	Type               uint8     // Network type
	EpochTime          time.Time // Network epoch time
	GenerationHashSeed []byte    // Network generation hash seed
}

func (network Network) AddTime(duration time.Duration) uint64 {
	currentTime := time.Now().UnixMilli()
	calcTime := int64(duration.Seconds()) - network.EpochTime.UnixMilli()

	return uint64(currentTime + calcTime)
}

func hexToBytes(hexString string) []byte {
	result, _ := hex.DecodeString(hexString)

	return result
}
