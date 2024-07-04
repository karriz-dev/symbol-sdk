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
	"github.com/stretchr/testify/assert"
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

func TestTransferTransactionV1(t *testing.T) {
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
		Deadline(time.Hour * 2).
		Signer(aliceKeyPair).
		TransferTransactionV1()

	serializeBytes, err := transferTx.
		Recipient(bobAddress).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).
		Message("Hello, Symbol - transact By Go SDK").Serialize()
	require.NoError(t, err)

	t.Log(common.BytesToHex(serializeBytes))

	assert.NoError(t, transferTx.Valid())
}

func TestTransferAliceToBob1XYM(t *testing.T) {
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
		Signer(aliceKeyPair).
		TransferTransactionV1()

	payloadBytes, err := transferTx.
		Recipient(bobAddress).
		Mosaics([]common.Mosaic{
			{
				MosaicId: 0x72C0212E67A08BCE,
				Amount:   1_000000,
			},
		}).
		Message("Hello, Symbol - transact By Go SDK").Sign()
	require.NoError(t, err)

	payload := "{\"payload\":\"" + common.BytesToHex(payloadBytes) + "\"}"

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
