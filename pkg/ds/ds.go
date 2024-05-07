package ds

import (
	"strconv"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
)

type DecentralizedStore interface {
	StoreDetailsOnChain(txs []byte, chainID int64, proposalID uint64) (string, error)
	QueryDetailsByTxID(txID string) ([]byte, error)
}
type ArWeave struct {
	Client *goar.Client
	Wallet *goar.Wallet
}

func NewArWeave(wallet *goar.Wallet, client *goar.Client) *ArWeave {
	return &ArWeave{
		Wallet: wallet,
		Client: client,
	}
}

func (ar *ArWeave) StoreDetailsOnChain(txs []byte, chainID int64, proposalID uint64) (string, error) {
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "title", Value: "b2-batch"},
		{Name: "chainID", Value: strconv.FormatInt(chainID, 10)},
		{Name: "ProposalID", Value: strconv.FormatUint(proposalID, 10)},
	}
	arTx, err := ar.Wallet.SendData(txs, tags)
	if err != nil {
		return "", err
	}
	return arTx.ID, nil
}

func (ar *ArWeave) QueryDetailsByTxID(txID string) ([]byte, error) {
	txs, err := ar.Client.GetTransactionData(txID, "json")
	if err != nil {
		return nil, err
	}
	return txs, nil
}

var _ DecentralizedStore = (*ArWeave)(nil)
