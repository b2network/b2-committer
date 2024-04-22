package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/b2network/b2committer/internal/types"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	url := "https://open-api-testnet.unisat.io/"
	client := NewBasicHTTPClient(url)
	headers := http.Header{}
	headers.Add("Authorization", "Bearer 4cb09e6eec9a0c3ebd07135b3817a3dcedcc9791d50fade4cb564e8ad68a7ac3")
	resp, err := client.Get(context.Background(), "/v1/indexer/tx/2f0025a6c5173a4cc5ea6c3eb85b5433c700dfeae34503394933712f8b3ffb26/outs", nil, headers)
	require.NoError(t, err)
	//b, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(b))

	var res types.APIBTCTxOutputs
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res.Data)
}

func TestContent(t *testing.T) {
	url := "https://open-api-testnet.unisat.io/"
	client := NewBasicHTTPClient(url)
	headers := http.Header{}
	headers.Add("Authorization", "Bearer 4cb09e6eec9a0c3ebd07135b3817a3dcedcc9791d50fade4cb564e8ad68a7ac3")
	resp, err := client.Get(context.Background(), "/v1/indexer/inscription/content/2f0025a6c5173a4cc5ea6c3eb85b5433c700dfeae34503394933712f8b3ffb26i0", nil, headers)
	require.NoError(t, err)
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))

	var res types.StateRootProposal
	err = json.Unmarshal(b, &res)
	require.NoError(t, err)
	fmt.Println(res)
}
