package b2node

import (
	"fmt"
	"github.com/b2network/b2committer/pkg/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

const (
	creatorAddress    = "0xb634434CA448c39b05b460dEC51f458EaC1e2759"
	creatorPrivateKey = "0a81baab0ca0b65d406d68c79945054b092cbe77499ca55c57b3ecfd33f1d551"
	contractAddress   = "0x85D40bDc724bcabF6D17d8343a74e0d916dfD40D"
	URL               = "https://habitat-hub-rpc.bsquared.network"
	chainID           = 1113
)

func TestQueryTimeoutPeriod(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	time, _ := committer.TimeoutPeriod(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	})
	fmt.Println(time)
}

func TestSetTimeoutPeriod(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.SetTimeoutPeriod(&bind.TransactOpts{
		From:   common.HexToAddress(creatorAddress),
		Signer: auth.Signer,
	}, big.NewInt(60*20))
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestAddChain(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.AddChain(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, 1113)
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestQueryAllChain(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	res, err := committer.AllChains(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	})
	fmt.Println(res)
	require.NoError(t, err)
	require.Containsf(t, res, uint64(1113), "error message %s")
}

func TestAddCommitter(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.AddProposer(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, common.HexToAddress(creatorAddress))
	require.NoError(t, err)
	require.NotEmpty(t, tx.Hash())
}

func TestQueryAllCommitter(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	res, err := committer.AllProposers(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	})
	fmt.Println(res)
	require.NoError(t, err)
	require.Containsf(t, res, common.HexToAddress(creatorAddress), "error message %s")
}

func TestRemoveCommitter(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.RemoveProposer(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, common.HexToAddress(creatorAddress))
	require.NotEmpty(t, tx.Hash())
}

func TestQueryLastProposalID(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	proposal, err := committer.GetLastProposal(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	}, chainID)
	require.NoError(t, err)
	fmt.Println(proposal)
}

func TestQueryProposalByID(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	proposal, err := committer.Proposal(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	}, chainID, 1)
	require.NoError(t, err)
	fmt.Println(proposal)
}

func TestSubmitProof(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.SubmitProof(&bind.TransactOpts{
		From:   common.HexToAddress(creatorAddress),
		Signer: auth.Signer,
	}, chainID, 1, "proofHash", "stateRoot", 1, 56908)
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestBitcoinTxHash(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.BitcoinTx(&bind.TransactOpts{
		From:   common.HexToAddress(creatorAddress),
		Signer: auth.Signer,
	}, chainID, 1, "txHash")
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestAr(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.ArweaveTx(&bind.TransactOpts{
		From:   common.HexToAddress(creatorAddress),
		Signer: auth.Signer,
	}, chainID, 1, "arTxhash")
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestIsProposalTimeout(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	res, err := committer.IsProposalTimeout(&bind.CallOpts{
		From: common.HexToAddress(creatorAddress),
	}, chainID, 2)
	fmt.Println(res)
}

func TestTimeoutProposal(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	require.NoError(t, err)
	privateKey, _ := crypto.HexToECDSA(creatorPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	require.NoError(t, err)
	tx, err := committer.TimeoutProposal(&bind.TransactOpts{
		From:   common.HexToAddress(creatorAddress),
		Signer: auth.Signer,
	}, chainID, 1)
	require.NoError(t, err)
	fmt.Println(tx.Hash())
}
