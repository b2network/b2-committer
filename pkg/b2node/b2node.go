package b2node

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/b2network/b2committer/pkg/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type NodeClient struct {
	PrivateKey      *ecdsa.PrivateKey
	ContractAddress string
	Address         string
	ChainID         int64
	Conn            *ethclient.Client
	Committer       *contract.Committer
	Auth            *bind.TransactOpts
}

func NewNodeClient(privateKeyStr string, chainID int64, address string, contractAddress string, conn *ethclient.Client) *NodeClient {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if nil != err {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	return &NodeClient{
		PrivateKey: privateKey,
		Address:    address,
		ChainID:    chainID,
		Conn:       conn,
		Committer:  committer,
		Auth:       auth,
	}
}

func (n NodeClient) AddCommitter(address string) (string, error) {
	tx, err := n.Committer.AddProposer(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, common.HexToAddress(address))
	if err != nil {
		return "", fmt.Errorf("[AddCommitter] err: %s", err)
	}
	return tx.Hash().String(), nil
}

func (n NodeClient) QueryAllProposers() ([]common.Address, error) {
	res, err := n.Committer.AllProposers(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("[QueryAllProposers] err: %s", err)
	}
	return res, nil
}

func (n NodeClient) RemoveProposers(address common.Address) (*types.Transaction, error) {
	tx, err := n.Committer.RemoveProposer(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, address)
	if err != nil {
		return nil, fmt.Errorf("[RemoveProposers] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) QueryLastProposal() (*contract.CommitterProposal, error) {
	proposal, err := n.Committer.GetLastProposal(&bind.CallOpts{
		From: common.HexToAddress(n.Address),
	}, uint64(n.ChainID))
	if err != nil {
		return nil, fmt.Errorf("[QueryLastProposalID] err: %s", err)
	}
	return &proposal, nil
}

func (n NodeClient) QueryProposalByID(id uint64) (*contract.CommitterProposal, error) {
	proposal, err := n.Committer.Proposal(&bind.CallOpts{
		From: common.HexToAddress(n.Address),
	}, uint64(n.ChainID), id)
	if err != nil {
		return nil, fmt.Errorf("[QueryProposalByID] err: %s", err)
	}
	return &proposal, nil
}

func (n NodeClient) SubmitProof(id uint64, proofHash string, stateRootHash string,
	startIndex uint64, endIndex uint64,
) (*types.Transaction, error) {
	tx, err := n.Committer.SubmitProof(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, uint64(n.ChainID), id, proofHash, stateRootHash, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) BitcoinTxHash(id uint64, txHash string) (*types.Transaction, error) {
	tx, err := n.Committer.BitcoinTx(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, uint64(n.ChainID), id, txHash)
	if err != nil {
		return nil, fmt.Errorf("[BitcoinTxHash] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) ArweaveTx(id uint64, txHash string) (*types.Transaction, error) {
	tx, err := n.Committer.ArweaveTx(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, uint64(n.ChainID), id, txHash)
	if err != nil {
		return nil, fmt.Errorf("[ArweaveTx] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) TimeoutProposal(id uint64) (*types.Transaction, error) {
	tx, err := n.Committer.TimeoutProposal(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, uint64(n.ChainID), id)
	if err != nil {
		return nil, fmt.Errorf("[TimeoutProposal] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) IsProposalTimeout(id uint64) (bool, error) {
	res, err := n.Committer.IsProposalTimeout(&bind.CallOpts{
		From: common.HexToAddress(n.Address),
	}, uint64(n.ChainID), id)
	if err != nil {
		return false, fmt.Errorf("[IsProposalTimeout] err: %s", err)
	}
	return res, nil
}
