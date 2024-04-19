package svc

import (
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/b2node"
	"github.com/b2network/b2committer/pkg/beacon"
	"github.com/b2network/b2committer/pkg/client"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/b2network/b2committer/pkg/ds"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/everFinance/goar"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/big"
	"time"
)

var svc *ServiceContext

type ServiceContext struct {
	Config                *types.Config
	RPC                   *ethclient.Client
	DB                    *gorm.DB
	BTCConfig             *types.BitcoinRPCConfig
	B2NodeConfig          *types.B2NODEConfig
	LatestBlockNumber     int64
	SyncedBlockNumber     int64
	SyncedBlockHash       common.Hash
	NodeClient            *b2node.NodeClient
	LatestBTCBlockNumber  int64
	BlobDataSource        *beacon.BlobDataSource
	SyncedBlobBlockNumber int64
	SyncedBlobBlockHash   common.Hash
	OpCommitterClient     *b2node.OpCommitterClient
	DecentralizedStore    ds.DecentralizedStore
}

func NewServiceContext(cfg *types.Config, bitcoinCfg *types.BitcoinRPCConfig, b2nodeConfig *types.B2NODEConfig) *ServiceContext {
	storage, err := gorm.Open(mysql.Open(cfg.MySQLDataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicf("[svc]gorm get db panic: %s\n", err)
	}

	sqlDB, err := storage.DB()
	if err != nil {
		log.Panicf("[svc]gorm get sqlDB panic: %s\n", err)
	}
	// SetMaxIdleConns
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
	// SetMaxOpenConns
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)
	// SetConnMaxLifetime
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MySQLConnMaxLifetime) * time.Second)

	rpc, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Panicf("[svc] get eth client panic: %s\n", err)
	}

	privateKeHex := b2nodeConfig.PrivateKey
	chainID := b2nodeConfig.ChainID
	address := b2nodeConfig.Address
	contractAddress := b2nodeConfig.CommitterAddress
	b2rpc, err := ethclient.Dial(b2nodeConfig.RPCUrl)
	if err != nil {
		log.Panicf("[svc] init b2node grpc panic: %s\n", err)
	}
	nodeClient := b2node.NewNodeClient(privateKeHex, chainID, address, contractAddress, b2rpc)

	l1Signer := ethTypes.NewCancunSigner(big.NewInt(chainID))
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(cfg.BeaconChainRPCUrl))
	l1BlobFetcher := sources.NewL1BeaconClient(l1Beacon, sources.L1BeaconClientConfig{FetchAllSidecars: false})
	bds := beacon.NewBlobDataSource(l1Signer, common.HexToAddress(cfg.BatcherInbox), common.HexToAddress(cfg.BatcherSender), l1BlobFetcher, rpc)

	proposer, err := op.NewProposer(common.HexToAddress(b2nodeConfig.OpProposersAddress), b2rpc)
	if err != nil {
		log.Panicf("[svc] init proposer contract panic: %s\n", err)
	}
	proposalManager, err := op.NewOpProposalManager(common.HexToAddress(b2nodeConfig.OpProposalManagerAddress), b2rpc)
	if err != nil {
		log.Panicf("[svc] init proposal manager contract panic: %s\n", err)
	}
	committer, err := op.NewOpCommitter(common.HexToAddress(b2nodeConfig.OpCommitterAddress), b2rpc)
	if err != nil {
		log.Panicf("[svc] init committer contract panic: %s\n", err)
	}
	opCommitterClient := b2node.NewOpCommitterClient(b2nodeConfig.PrivateKey, b2nodeConfig.ChainID, proposer, committer, proposalManager)

	svc = &ServiceContext{
		BTCConfig:         bitcoinCfg,
		DB:                storage,
		Config:            cfg,
		RPC:               rpc,
		LatestBlockNumber: cfg.InitBlockNumber,
		B2NodeConfig:      b2nodeConfig,
		NodeClient:        nodeClient,
		BlobDataSource:    bds,
		OpCommitterClient: opCommitterClient,
	}

	dsType := cfg.DSType
	if dsType == "" {
		panic("Invalid dsType")
	}
	if dsType == "arweave" {
		if cfg.ArweaveRPC == "" {
			panic("Invalid arweave rpc")
		}
		if cfg.ArweaveWallet == "" {
			panic("Invalid arweaveWallet path")
		}
		arClient := goar.NewClient(cfg.ArweaveRPC)
		wallet := cfg.ArweaveWallet
		arNode := cfg.ArweaveRPC
		w, err := goar.NewWalletFromPath(wallet, arNode)
		if err != nil {
			log.Panicf("[svc] init arweave wallet panic: %s\n", err)
		}

		svc.DecentralizedStore = ds.NewArWeave(w, arClient)
	}
	return svc
}
