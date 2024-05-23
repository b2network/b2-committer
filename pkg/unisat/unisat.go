package unisat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/client"
)

type UstHTTPClient struct {
	auth string
	cl   client.HTTP
}

func NewUnisatHTTPClient(cl client.HTTP, auth string) *UstHTTPClient {
	return &UstHTTPClient{auth, cl}
}

func (cl *UstHTTPClient) apiReq(ctx context.Context, dest any, reqPath string, reqQuery url.Values) error {
	headers := http.Header{}
	headers.Add("Accept", "application/json")
	headers.Add("Authorization", "Bearer "+cl.auth)
	resp, err := cl.cl.Get(ctx, reqPath, reqQuery, headers)
	if err != nil {
		return fmt.Errorf("http Get failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		errMsg, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return fmt.Errorf("failed request with status %d: %s", resp.StatusCode, string(errMsg))
	}
	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		_ = resp.Body.Close()
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}
	return nil
}

func (cl *UstHTTPClient) QueryAPIBTCTxOutputsByTxID(ctx context.Context, txID string) (*types.APIBTCTxOutputs, error) {
	var res types.APIBTCTxOutputs
	err := cl.apiReq(ctx, &res, "/v1/indexer/tx/"+txID+"/outs", nil)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (cl *UstHTTPClient) QueryStateRootProposalByInsID(ctx context.Context, insID string) (*types.BtcStateRootProposal, error) {
	var res types.BtcStateRootProposal
	err := cl.apiReq(ctx, &res, "/v1/indexer/inscription/content/"+insID, nil)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
