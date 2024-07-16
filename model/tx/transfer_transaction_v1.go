package tx

import (
	"github.com/karriz-dev/symbol-sdk/errors"
	"github.com/karriz-dev/symbol-sdk/model/account"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/message"
	"github.com/karriz-dev/symbol-sdk/model/mosaic"
	"github.com/karriz-dev/symbol-sdk/network"
)

type TransferTransactionV1 struct {
	BaseTransaction

	recipient account.Address
	mosaics   []mosaic.Mosaic
	message   message.Message

	mosaicsCount decimal.UInt8

	transferTransactionBodyReserved_1 decimal.UInt8
	transferTransactionBodyReserved_2 decimal.UInt32
}

func NewTransferTransactionV1(network network.Network, maxFee decimal.UInt64, deadline decimal.UInt64, signer account.PublicKey, isEmbedded bool) TransferTransactionV1 {
	tx := BaseTransaction{
		size:                            decimal.NewUInt32(32),
		version:                         decimal.NewUInt8(1),
		network:                         network,
		txType:                          decimal.NewUInt16(0x4154),
		fee:                             maxFee,
		deadline:                        deadline,
		verifiableEntityHeaderReserved1: decimal.NewUInt32(0),
		entityBodyReserved1:             decimal.NewUInt32(0),
		signer:                          signer,
		isEmbedded:                      isEmbedded,
	}

	return TransferTransactionV1{
		BaseTransaction:                   tx,
		transferTransactionBodyReserved_1: decimal.NewUInt8(0),
		transferTransactionBodyReserved_2: decimal.NewUInt32(0),
	}
}

func (tx *TransferTransactionV1) Recipient(recipient account.Address) *TransferTransactionV1 {
	tx.recipient = recipient

	return tx
}

func (tx *TransferTransactionV1) Mosaics(mosaics []mosaic.Mosaic) *TransferTransactionV1 {
	mosaicsLength := len(mosaics)
	if mosaicsLength > 255 {
		// TODO :: need error
		return tx
	}

	tx.mosaics = mosaics
	tx.mosaicsCount = decimal.NewUInt8(uint8(mosaicsLength))

	tx.size.Add(uint32(mosaicsLength * 16))

	return tx
}

func (tx *TransferTransactionV1) Message(msg string) *TransferTransactionV1 {
	// TODO :: need encrypt message
	if len(msg) > 65535 {
		return tx
	}

	tx.message = message.Message(msg)

	if len(msg) > 0 {
		tx.size.Add(uint32(len(msg) + 1))
	}

	return tx
}

func (tx TransferTransactionV1) Serialize() ([]byte, error) {
	// serialize base tx attrs
	serializeData, err := tx.BaseTransaction.Serialize()
	if err != nil {
		return nil, err
	}
	if len(tx.recipient) == 0 {
		return nil, errors.ErrRecipientNotValid
	}

	// serialize attrs
	serializeData = append(serializeData, tx.recipient[:]...)
	serializeData = append(serializeData, tx.message.LenBytes()...)
	serializeData = append(serializeData, tx.mosaicsCount.Bytes()...)
	serializeData = append(serializeData, tx.transferTransactionBodyReserved_1.Bytes()...)
	serializeData = append(serializeData, tx.transferTransactionBodyReserved_2.Bytes()...)

	// serialize mosiacs
	for _, mosaic := range tx.mosaics {
		serializeData = append(serializeData, mosaic.Bytes()...)
	}

	// serialize message
	serializeData = append(serializeData, tx.message.Bytes()...)

	return serializeData, nil
}
