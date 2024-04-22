package unisat

import (
	"context"
	"fmt"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/client"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryInscribeNumber(t *testing.T) {
	url := "https://open-api-testnet.unisat.io/"
	unisat := NewUnisatHTTPClient(client.NewBasicHTTPClient(url), "4cb09e6eec9a0c3ebd07135b3817a3dcedcc9791d50fade4cb564e8ad68a7ac3")
	var res types.APIBTCTxOutputs
	err := unisat.apiReq(context.Background(), &res, "/v1/indexer/tx/2f0025a6c5173a4cc5ea6c3eb85b5433c700dfeae34503394933712f8b3ffb26/outs", nil)
	require.NoError(t, err)
	fmt.Println(res)
	fmt.Println(res.Data[0].Inscriptions[0].InscriptionID)
}

func TestQueryContent(t *testing.T) {
	url := "https://open-api-testnet.unisat.io/"
	unisat := NewUnisatHTTPClient(client.NewBasicHTTPClient(url), "4cb09e6eec9a0c3ebd07135b3817a3dcedcc9791d50fade4cb564e8ad68a7ac3")
	var res types.StateRootProposal
	err := unisat.apiReq(context.Background(), &res, "/v1/indexer/inscription/content/2f0025a6c5173a4cc5ea6c3eb85b5433c700dfeae34503394933712f8b3ffb26i0", nil)
	require.NoError(t, err)
	fmt.Println(res)
}
