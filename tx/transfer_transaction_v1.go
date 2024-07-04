package tx

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha512"

	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/types"
)

type TransferTransactionV1 struct {
	Transaction
	innerSerializeBytes []byte // inner serialize

	recipient common.Address
	mosaics   []common.Mosaic
	message   common.Message

	messageLength common.MessageLength
	mosaicsCount  common.MosaicsCount

	transferTransactionBodyReserved_1 byte
	transferTransactionBodyReserved_2 [4]byte
}

func (transactionFactory *TransactionFactory) TransferTransactionV1() TransferTransactionV1 {
	commonTx := Transaction{
		size:                            128, // 128 = common tx size
		version:                         0x01,
		network:                         transactionFactory.network,
		txType:                          0x4154, // transfer_transaction_v1
		fee:                             transactionFactory.maxFee,
		deadline:                        transactionFactory.deadline,
		verifiableEntityHeaderReserved1: []byte{0x00, 0x00, 0x00, 0x00},
		entityBodyReserved1:             []byte{0x00, 0x00, 0x00, 0x00},
		signer:                          transactionFactory.signer,
	}

	// TODO:: needs error check
	innerSerializeData, _ := commonTx.Serialize()

	commonTx.size += 24 // recipient
	commonTx.size += 2  // message length
	commonTx.size += 1  // mosaic size
	commonTx.size += 1  // transferTransactionBodyReserved_1
	commonTx.size += 4  // transferTransactionBodyReserved_2

	return TransferTransactionV1{
		Transaction:                       commonTx,
		innerSerializeBytes:               innerSerializeData,
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

func (transferTransactionV1 *TransferTransactionV1) Serialize() ([]byte, error) {
	// serialize inner common tx attrs
	serializeData, err := transferTransactionV1.Transaction.Serialize()
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

func (transferTransactionV1 *TransferTransactionV1) Sign() ([]byte, error) {
	data, err := transferTransactionV1.Serialize()
	if err != nil {
		return nil, err
	}

	appendedData := append(transferTransactionV1.network.GenerationHashSeed, data...)

	hasher := sha512.New()
	hasher.Write(appendedData)
	hashedData := hasher.Sum(nil)

	edPrivateKey := ed25519.NewKeyFromSeed(transferTransactionV1.signer.PrivateKey[:])
	sign, err := edPrivateKey.Sign(nil, hashedData,
		&ed25519.Options{
			Hash: crypto.SHA512,
		},
	)
	if err != nil {
		return nil, err
	}

	transferTransactionV1.signature = common.Signature(sign)
	appendedSerializedData, err := transferTransactionV1.Serialize()
	if err != nil {
		return nil, err
	}

	return appendedSerializedData, nil
}

func (transferTransactionV1 *TransferTransactionV1) Valid() error {
	data, err := transferTransactionV1.Serialize()
	if err != nil {
		return err
	}

	appendedData := append(transferTransactionV1.network.GenerationHashSeed, data...)

	hasher := sha512.New()
	hasher.Write(appendedData)
	hashedData := hasher.Sum(nil)

	edPrivateKey := ed25519.NewKeyFromSeed(transferTransactionV1.signer.PrivateKey[:])

	sign, err := edPrivateKey.Sign(nil, hashedData,
		&ed25519.Options{
			Hash: crypto.SHA512,
		},
	)
	if err != nil {
		return err
	}

	err = ed25519.VerifyWithOptions(edPrivateKey.Public().(ed25519.PublicKey), hashedData, sign,
		&ed25519.Options{
			Hash: crypto.SHA512,
		})

	return err
}

// func (transferTransactionV1 *TransferTransactionV1) Valid() error {
// 	return nil
// }

// func (transferTransactionV1 TransferTransactionV1) Hash() common.Hash {
// 	return common.Hash{}
// }
