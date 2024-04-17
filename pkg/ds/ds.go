package ds

import (
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"strconv"
)

type DecentralizedStore interface {
	StoreTxsOnChain(txs []byte, chainID int64, proposalID uint64) (string, error)

	QueryTxsByTxID(txID string) ([]byte, error)
}
type ArWeave struct {
	URL        string //rpc url
	WalletPath string
	Client     *goar.Client
}

func NewArWeave(url, walletPath string, client *goar.Client) *ArWeave {
	return &ArWeave{
		URL:        url,
		WalletPath: walletPath,
		Client:     client,
	}
}

func (ar *ArWeave) StoreTxsOnChain(txs []byte, chainID int64, proposalID uint64) (string, error) {
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "title", Value: "b2-batch"},
		{Name: "chainID", Value: strconv.FormatInt(chainID, 10)},
		{Name: "ProposalID", Value: strconv.FormatUint(proposalID, 10)},
	}
	w, err := goar.NewWalletFromPath(ar.WalletPath, ar.URL)
	if err != nil {
		return "", err
	}
	arTx, err := w.SendData(txs, tags)
	if err != nil {
		return "", err
	}
	return arTx.ID, nil
}

func (ar *ArWeave) QueryTxsByTxID(txID string) ([]byte, error) {
	txs, err := ar.Client.GetTransactionData(txID, "json")
	if err != nil {
		return nil, err
	}
	return txs, nil
}

var _ DecentralizedStore = (*ArWeave)(nil)
