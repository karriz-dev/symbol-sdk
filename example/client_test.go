package example

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	symbolsdk "github.com/karriz-dev/symbol-sdk"
	"github.com/karriz-dev/symbol-sdk/common"
	"github.com/karriz-dev/symbol-sdk/network"
	"github.com/stretchr/testify/require"
)

const (
	SymbolTestNetworkUrl = "https://001-sai-dual.symboltest.net:3001"
)

var TestAddressList []string

func init() {
	data, err := os.ReadFile("./address.txt")
	if err != nil {
		panic(err)
	}

	TestAddressList = strings.Split(string(data), "\n")
}

func TestGetNetworkInfo(t *testing.T) {
	symbolRestClient := http.Client{
		Timeout: time.Second * 3,
	}

	response, err := symbolRestClient.Get(SymbolTestNetworkUrl + "/network/properties")
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	m, _ := json.MarshalIndent(result, "", "\t")
	t.Log(string(m))

}

func TestTransactionSignAndVerify(t *testing.T) {
	require.Len(t, TestAddressList, 2)

	aliceKeyPair, err := common.HexToKeyPair(TestAddressList[0])
	require.NoError(t, err)

	bobKeyPair, err := common.HexToKeyPair(TestAddressList[1])
	require.NoError(t, err)

	aliceAddress := common.PublicKeyToAddress(aliceKeyPair.PublicKey, network.TESTNET)
	bobAddress := common.PublicKeyToAddress(bobKeyPair.PublicKey, network.TESTNET)

	// NOTE :: if you change address.txt on yours edit this expected
	require.Equal(t, "TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI", aliceAddress.Encode())
	require.Equal(t, "TBROSRKD5LONZYOP4II7JJTLS5IY6IOOG34DZHI", bobAddress.Encode())

	facade := symbolsdk.NewSymbolFacade("testnet")
	transferTx := facade.TransactionFactory.
		MaxFee(1_000000).
		Deadline(time.Minute * 10).
		Signer(aliceKeyPair.PublicKey).
		TransferTransactionV1()

	transferTx.
		Recipient(bobAddress).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).
		Message("Hello, Symbol - transact By Go SDK")

	signature, err := facade.TransactionFactory.Sign(&transferTx, aliceKeyPair.PrivateKey)
	require.NoError(t, err)

	transactionBytes, err := transferTx.Serialize()
	require.NoError(t, err)

	verify := facade.TransactionFactory.Verify(transactionBytes, signature[:], aliceKeyPair.PublicKey)
	require.NoError(t, verify)

	transferTx.AttachSignature(signature)
	signedTxPayload, err := transferTx.Serialize()
	require.NoError(t, err)

	t.Log(common.BytesToJSONPayload(signedTxPayload))
}

func TestTransactionNetworkAnnounce(t *testing.T) {
	require.Len(t, TestAddressList, 2)

	aliceKeyPair, err := common.HexToKeyPair(TestAddressList[0])
	require.NoError(t, err)

	bobKeyPair, err := common.HexToKeyPair(TestAddressList[1])
	require.NoError(t, err)

	aliceAddress := common.PublicKeyToAddress(aliceKeyPair.PublicKey, network.TESTNET)
	bobAddress := common.PublicKeyToAddress(bobKeyPair.PublicKey, network.TESTNET)

	// NOTE :: if you change address.txt on yours edit this expected
	require.Equal(t, "TDWNFH2JA5FG3L5LTGYIS5TB475TENZVYJ4CQCI", aliceAddress.Encode())
	require.Equal(t, "TBROSRKD5LONZYOP4II7JJTLS5IY6IOOG34DZHI", bobAddress.Encode())

	facade := symbolsdk.NewSymbolFacade("testnet")
	transferTx := facade.TransactionFactory.
		MaxFee(1_000000).
		Deadline(time.Minute * 10).
		Signer(aliceKeyPair.PublicKey).
		TransferTransactionV1()

	transferTx.
		Recipient(bobAddress).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).
		Message("Hello, Symbol - transact By Go SDK")

	signature, err := facade.TransactionFactory.Sign(&transferTx, aliceKeyPair.PrivateKey)
	require.NoError(t, err)

	transferTx.AttachSignature(signature)
	signedTxPayload, err := transferTx.Serialize()
	require.NoError(t, err)

	payload := common.BytesToJSONPayload(signedTxPayload)

	t.Log(payload)

	symbolRestClient := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodPut, SymbolTestNetworkUrl+"/transactions", bytes.NewBufferString(payload))
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
