package svc

import (
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/b2node"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var svc *ServiceContext

type ServiceContext struct {
	Config               *types.Config
	RPC                  *ethclient.Client
	DB                   *gorm.DB
	BTCConfig            *types.BitcoinRPCConfig
	B2NodeConfig         *types.B2NODEConfig
	LatestBlockNumber    int64
	SyncedBlockNumber    int64
	SyncedBlockHash      common.Hash
	NodeClient           *b2node.NodeClient
	LatestBTCBlockNumber int64
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

	svc = &ServiceContext{
		BTCConfig:         bitcoinCfg,
		DB:                storage,
		Config:            cfg,
		RPC:               rpc,
		LatestBlockNumber: cfg.InitBlockNumber,
		B2NodeConfig:      b2nodeConfig,
		NodeClient:        nodeClient,
	}
	return svc
}
