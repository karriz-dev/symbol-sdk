package tx

import (
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/types"
)

type TransferTransactionV1 struct {
	Transaction

	recipient common.Address
	mosaics   []common.Mosaic
	message   common.Message

	messageLength common.MessageLength
	mosaicsCount  common.MosaicsCount

	transferTransactionBodyReserved_1 byte
	transferTransactionBodyReserved_2 [4]byte
}

type EmbeddedTransferTransactionV1 struct {
	EmbeddedTransaction

	recipient common.Address
	mosaics   []common.Mosaic
	message   common.Message

	messageLength common.MessageLength
	mosaicsCount  common.MosaicsCount

	transferTransactionBodyReserved_1 byte
	transferTransactionBodyReserved_2 [4]byte
}

func (transactionFactory *TransactionFactory) TransferTransactionV1() TransferTransactionV1 {
	txHeader := Transaction{
		size:                            128, // 128 = tx header default size
		version:                         0x01,
		network:                         transactionFactory.network,
		txType:                          0x4154, // transfer_transaction_v1
		fee:                             transactionFactory.maxFee,
		deadline:                        transactionFactory.deadline,
		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
		signer:                          transactionFactory.signer,
	}

	// adding TransferTransactionV1 entity size
	txHeader.size += 32

	return TransferTransactionV1{
		Transaction:                       txHeader,
		transferTransactionBodyReserved_1: 0x00,
		transferTransactionBodyReserved_2: [4]byte{0x00, 0x00, 0x00, 0x00},
	}
}

func (embeddedTransactionFactory *EmbeddedTransactionFactory) EmbeddedTransferTransactionV1() EmbeddedTransferTransactionV1 {
	txHeader := EmbeddedTransaction{
		size:                48, // 48 = embedded tx header default size
		version:             0x01,
		network:             embeddedTransactionFactory.network,
		txType:              0x4154, // transfer_transaction_v1
		entityBodyReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		signer:              embeddedTransactionFactory.signer,
	}

	// adding TransferTransactionV1 entity size
	txHeader.size += 32

	return EmbeddedTransferTransactionV1{
		EmbeddedTransaction:               txHeader,
		transferTransactionBodyReserved_1: 0x00,
		transferTransactionBodyReserved_2: [4]byte{0x00, 0x00, 0x00, 0x00},
	}
}

func (transferTransactionV1 *TransferTransactionV1) Recipient(recipient common.Address) *TransferTransactionV1 {
	transferTransactionV1.recipient = recipient

	return transferTransactionV1
}

func (transferTransactionV1 *TransferTransactionV1) Mosaics(mosaics []common.Mosaic) *TransferTransactionV1 {
	transferTransactionV1.mosaics = mosaics
	transferTransactionV1.mosaicsCount = common.MosaicsCount(len(mosaics))

	transferTransactionV1.size += types.TransactionSize(len(mosaics) * 16) // mosaic length

	return transferTransactionV1
}

func (transferTransactionV1 *TransferTransactionV1) Message(message string) *TransferTransactionV1 {
	transferTransactionV1.message = common.Message(message)
	transferTransactionV1.messageLength = common.MessageLength(len(message))

	transferTransactionV1.size += types.TransactionSize(len(message) + 1) // message length

	return transferTransactionV1
}

func (transferTransactionV1 TransferTransactionV1) Serialize() ([]byte, error) {
	// serialize inner common tx attrs
	serializeData, err := transferTransactionV1.Transaction.serialize()
	if err != nil {
		return nil, err
	}

	// serialize attrs
	serializeData = append(serializeData, transferTransactionV1.recipient.Bytes()...)
	serializeData = append(serializeData, transferTransactionV1.messageLength.Bytes()...)
	serializeData = append(serializeData, transferTransactionV1.mosaicsCount.Byte())
	serializeData = append(serializeData, transferTransactionV1.transferTransactionBodyReserved_1)
	serializeData = append(serializeData, transferTransactionV1.transferTransactionBodyReserved_2[:]...)

	// serialize mosiacs
	for _, mosaic := range transferTransactionV1.mosaics {
		serializeData = append(serializeData, mosaic.Bytes()...)
	}

	// serialize message
	serializeData = append(serializeData, transferTransactionV1.message.Bytes()...)

	return serializeData, nil
}

func (embeddedTransferTransactionV1 *EmbeddedTransferTransactionV1) Recipient(recipient common.Address) *EmbeddedTransferTransactionV1 {
	embeddedTransferTransactionV1.recipient = recipient

	return embeddedTransferTransactionV1
}

func (embeddedTransferTransactionV1 *EmbeddedTransferTransactionV1) Mosaics(mosaics []common.Mosaic) *EmbeddedTransferTransactionV1 {
	embeddedTransferTransactionV1.mosaics = mosaics
	embeddedTransferTransactionV1.mosaicsCount = common.MosaicsCount(len(mosaics))

	embeddedTransferTransactionV1.size += types.TransactionSize(len(mosaics) * 16) // mosaic length

	return embeddedTransferTransactionV1
}

func (embeddedTransferTransactionV1 *EmbeddedTransferTransactionV1) Message(message string) *EmbeddedTransferTransactionV1 {
	embeddedTransferTransactionV1.message = common.Message(message)
	embeddedTransferTransactionV1.messageLength = common.MessageLength(len(message))

	embeddedTransferTransactionV1.size += types.TransactionSize(len(message) + 1) // message length

	return embeddedTransferTransactionV1
}

func (embeddedTransferTransactionV1 EmbeddedTransferTransactionV1) Serialize() ([]byte, error) {
	// serialize inner common tx attrs
	serializeData, err := embeddedTransferTransactionV1.EmbeddedTransaction.serialize()
	if err != nil {
		return nil, err
	}

	// serialize attrs
	serializeData = append(serializeData, embeddedTransferTransactionV1.recipient.Bytes()...)
	serializeData = append(serializeData, embeddedTransferTransactionV1.messageLength.Bytes()...)
	serializeData = append(serializeData, embeddedTransferTransactionV1.mosaicsCount.Byte())
	serializeData = append(serializeData, embeddedTransferTransactionV1.transferTransactionBodyReserved_1)
	serializeData = append(serializeData, embeddedTransferTransactionV1.transferTransactionBodyReserved_2[:]...)

	// serialize mosiacs
	for _, mosaic := range embeddedTransferTransactionV1.mosaics {
		serializeData = append(serializeData, mosaic.Bytes()...)
	}

	// serialize message
	serializeData = append(serializeData, embeddedTransferTransactionV1.message.Bytes()...)

	return serializeData, nil
}
