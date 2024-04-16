package b2node

import (
	"fmt"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	WalletAddress            = "0xb634434CA448c39b05b460dEC51f458EaC1e2759"
	WalletPrivateKey         = "0a81baab0ca0b65d406d68c79945054b092cbe77499ca55c57b3ecfd33f1d551"
	Sepolia_URL              = "https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30/"
	Sepolia_chainID          = 11155111
	proposerContractAddress  = "0x6bA5d52CA1B931E89611995F53dde2E1f914b0a7"
	proposalManagerContract  = "0x1A9Df4e0949a8699739B63F0155F41e0505df593"
	committerContractAddress = "0x0484fE0888684b39B2cce0a1644e44dDFCAA0054"
)

var opCommitterClient *OpCommitterClient

func setup(t *testing.T) {
	conn, err := ethclient.Dial(Sepolia_URL)
	require.NoError(t, err)
	defer conn.Close()

	proposer, err := op.NewProposer(common.HexToAddress(proposerContractAddress), conn)
	require.NoError(t, err)
	proposalManager, err := op.NewOpProposalManager(common.HexToAddress(proposalManagerContract), conn)
	require.NoError(t, err)
	committer, err := op.NewOpCommitter(common.HexToAddress(committerContractAddress), conn)
	require.NoError(t, err)
	opCommitterClient = NewOpCommitterClient(WalletPrivateKey, Sepolia_chainID, proposer, committer, proposalManager)
}

func TestQueryAllProposals(t *testing.T) {
	setup(t)
	res, err := opCommitterClient.QueryAllProposers()
	require.NoError(t, err)
	fmt.Println(res)
}

func TestAddProposer(t *testing.T) {
	setup(t)
	tx, err := opCommitterClient.AddProposer(common.HexToAddress(WalletAddress))
	require.NoError(t, err)
	fmt.Println(tx)
}

func TestIsProposer(t *testing.T) {
	setup(t)
	res, err := opCommitterClient.Proposer.IsProposer(&bind.CallOpts{}, common.HexToAddress(WalletAddress))
	require.NoError(t, err)
	fmt.Println(res)
}

func TestGetLastTxsRootProposal(t *testing.T) {
	setup(t)
	proposal, err := opCommitterClient.ProposalManager.GetLastTxsRootProposal(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Println(proposal)
}
