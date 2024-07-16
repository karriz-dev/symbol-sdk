package example_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/karriz-dev/symbol-sdk/example"
	"github.com/stretchr/testify/require"
)

func TestGetNetworkRecommendedFee(t *testing.T) {
	symbolRestClient := http.Client{
		Timeout: time.Second * 3,
	}

	response, err := symbolRestClient.Get(example.SymbolTestNetworkUrl + "/network/fees/transaction")
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	m, _ := json.MarshalIndent(result, "", "\t")
	t.Log(string(m))
}

func TestGetNetworkInfo(t *testing.T) {
	symbolRestClient := http.Client{
		Timeout: time.Second * 3,
	}

	response, err := symbolRestClient.Get(example.SymbolTestNetworkUrl + "/network/properties")
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	m, _ := json.MarshalIndent(result, "", "\t")
	t.Log(string(m))
}
