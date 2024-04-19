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
	WalletAddress = "0xb634434CA448c39b05b460dEC51f458EaC1e2759"

	WalletPrivateKey         = "0a81baab0ca0b65d406d68c79945054b092cbe77499ca55c57b3ecfd33f1d551"
	Sepolia_URL              = "https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30/"
	Sepolia_chainID          = 11155111
	proposerContractAddress  = "0x8C1ceB039414204809a343C1145a26E4a927Ac0f"
	proposalManagerContract  = "0x18465f258b51CfB576564791016152F583d1acD1"
	committerContractAddress = "0x2D3F3B68eDb74a9665aC5394eC51881d54529aa5"
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

func TestSetCommitterContract(t *testing.T) {
	setup(t)
	tx, err := opCommitterClient.SetCommitter(committerContractAddress)
	require.NoError(t, err)
	fmt.Println(tx)
}

func TestGetLastTxsRootProposal(t *testing.T) {
	setup(t)
	proposal, err := opCommitterClient.ProposalManager.GetLastTxsRootProposal(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Println(proposal)
}

func TestGetTxsRootProposalByID(t *testing.T) {
	setup(t)
	proposal1, err := opCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, 1)
	require.NoError(t, err)
	fmt.Println(proposal1)
	proposal2, err := opCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, 2)
	require.NoError(t, err)
	fmt.Println(proposal2)
	proposal3, err := opCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, 3)
	require.NoError(t, err)
	fmt.Println(proposal3)
}

func TestQueryCommitterProposer(t *testing.T) {
	setup(t)
	proposer, err := opCommitterClient.Committer.Proposer(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Println(proposer)
}
