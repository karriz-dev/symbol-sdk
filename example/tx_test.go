package example_test

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	symbolsdk "github.com/karriz-dev/symbol-sdk"
	"github.com/karriz-dev/symbol-sdk/example"
	"github.com/karriz-dev/symbol-sdk/model/decimal"
	"github.com/karriz-dev/symbol-sdk/model/mosaic"
	"github.com/karriz-dev/symbol-sdk/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionSignAndVerify(t *testing.T) {
	// NOTE :: if you change address.txt on yours edit this expected
	require.Equal(t, "TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI", example.AliceAccount.EncodedAddress())
	require.Equal(t, "TBROSRKD5LONZYOP4II7JJTLS5IY6IOOG34DZHI", example.BobAccount.EncodedAddress())
	require.Equal(t, "TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA", example.ServiceProviderAccount.EncodedAddress())

	t.Log("example.BobAccount.Address", example.BobAccount.Address[:])

	facade := symbolsdk.NewSymbolFacade("testnet")
	transferTx := facade.TransactionFactory.
		MaxFee(10000).
		Deadline(time.Hour * 2).
		Signer(example.AliceAccount.PublicKey).
		TransferTransactionV1(false)

	transferTx.
		Recipient(example.BobAccount.Address).
		Mosaics([]mosaic.Mosaic{
			{
				MosaicId: decimal.NewUInt64(0x72C0212E67A08BCE),
				Amount:   decimal.NewUInt64(1_000000),
			},
		}).
		Message("Hello, Symbol - transact By Go SDK")

	signature, err := facade.TransactionFactory.Sign(transferTx, example.AliceAccount.PrivateKey)
	require.NoError(t, err)

	transactionBytes, err := transferTx.Serialize()
	require.NoError(t, err)

	verify := facade.TransactionFactory.Verify(transactionBytes, signature[:], example.AliceAccount.PublicKey)
	require.NoError(t, verify)

	transferTx.AttachSignature(signature)

	signedTxPayload, err := transferTx.Serialize()
	require.NoError(t, err)

	// tx size check
	txSize := binary.LittleEndian.Uint16(append([]byte{signedTxPayload[0]}, signedTxPayload[1]))
	assert.Equal(t, uint16(len(signedTxPayload)), txSize)

	t.Log(signedTxPayload)
}

func TestTransactionNetworkAnnounce(t *testing.T) {
	// NOTE :: if you change address.txt on yours edit this expected
	require.Equal(t, "TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI", example.AliceAccount.EncodedAddress())
	require.Equal(t, "TBROSRKD5LONZYOP4II7JJTLS5IY6IOOG34DZHI", example.BobAccount.EncodedAddress())
	require.Equal(t, "TBLWGZ5W6VYS7BAE3O6VMN5VIW4FTC3BCDEYDMA", example.ServiceProviderAccount.EncodedAddress())

	facade := symbolsdk.NewSymbolFacade("testnet")
	transferTx := facade.TransactionFactory.
		MaxFee(10000).
		Deadline(time.Minute * 1).
		Signer(example.AliceAccount.PublicKey).
		TransferTransactionV1(false)

	transferTx.
		Recipient(example.BobAccount.Address).
		Mosaics([]mosaic.Mosaic{
			{
				MosaicId: decimal.NewUInt64(0x72C0212E67A08BCE),
				Amount:   decimal.NewUInt64(1_000000),
			},
		}).
		Message("Hello, Symbol - transact By Go SDK")

	signature, err := facade.TransactionFactory.Sign(transferTx, example.AliceAccount.PrivateKey)
	require.NoError(t, err)

	transactionBytes, err := transferTx.Serialize()
	require.NoError(t, err)

	verify := facade.TransactionFactory.Verify(transactionBytes, signature[:], example.AliceAccount.PublicKey)
	require.NoError(t, verify)

	transferTx.AttachSignature(signature)

	signedTxPayload, err := transferTx.Serialize()
	require.NoError(t, err)

	// tx size check
	txSize := binary.LittleEndian.Uint16(append([]byte{signedTxPayload[0]}, signedTxPayload[1]))
	assert.Equal(t, uint16(len(signedTxPayload)), txSize)

	payload := util.BytesToJSONPayload(signedTxPayload)

	t.Log(payload)

	symbolRestClient := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodPut, example.SymbolTestNetworkUrl+"/transactions", bytes.NewBufferString(payload))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	response, err := symbolRestClient.Do(req)
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	m, _ := json.MarshalIndent(result, "", "\t")
	t.Log(string(m))
}
