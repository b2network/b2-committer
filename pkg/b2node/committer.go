package b2node

import (
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type OpCommitterClient struct {
	Proposer        *op.Proposer
	Committer       *op.OpCommitter
	ProposalManager *op.OpProposalManager
	Auth            *bind.TransactOpts
}

func NewOpCommitterClient(privateKeyStr string, chainID int64, proposer *op.Proposer, committer *op.OpCommitter, proposalManager *op.OpProposalManager) *OpCommitterClient {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if nil != err {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	return &OpCommitterClient{
		Proposer:        proposer,
		Committer:       committer,
		ProposalManager: proposalManager,
		Auth:            auth,
	}
}

func (client *OpCommitterClient) QueryAllProposers() ([]common.Address, error) {
	res, err := client.Proposer.AllProposers(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (client *OpCommitterClient) AddProposer(address common.Address) (*ethTypes.Transaction, error) {
	tx, err := client.Proposer.AddProposer(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, address)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (client *OpCommitterClient) SubmitTxsRoot(proposal *types.TxsRootProposal) (*ethTypes.Transaction, error) {
	tx, err := client.Committer.SubmitTxsRoot(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, proposal.ProposalID, proposal.StartTimestamp, proposal.EndTimestamp, proposal.StartBlockNumber, proposal.EndBlockNumber, proposal.TxsRoot, proposal.BlockList)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (client *OpCommitterClient) SubmitStateRoot(proposal *types.StateRootProposal) (*ethTypes.Transaction, error) {
	tx, err := client.Committer.SubmitStateRoot(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, proposal.ProposalID, proposal.OutputRoot, proposal.StartL1Timestamp, proposal.EndL1Timestamp, proposal.StartL2BlockNumber,
		proposal.EndL2BlockNumber, proposal.OutputStartIndex, proposal.OutputEndIndex)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (client *OpCommitterClient) DsHash(proposalID uint64, proposalType uint64, dsType uint64, txsHash string) (*ethTypes.Transaction, error) {
	tx, err := client.Committer.DsTxHash(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, proposalID, proposalType, dsType, txsHash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (client *OpCommitterClient) BitcoinTxHash(proposalID uint64, btcTxHash string) (*ethTypes.Transaction, error) {
	tx, err := client.Committer.BitcoinTxHash(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, proposalID, btcTxHash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (client *OpCommitterClient) SetCommitter(committerContractAddress string) (*ethTypes.Transaction, error) {
	tx, err := client.ProposalManager.SetCommitter(&bind.TransactOpts{
		From:   client.Auth.From,
		Signer: client.Auth.Signer,
	}, common.HexToAddress(committerContractAddress))
	if err != nil {
		return nil, err
	}
	return tx, nil
}
