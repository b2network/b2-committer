package types

import (
	"github.com/b2network/b2committer/pkg/log"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat                   string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource             string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3366)/b2_committer_main?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns           int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns           int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime        int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl                      string `env:"RPC_URL" envDefault:"https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30"`
	BeaconChainID               int64  `env:"BEACON_CHAIN_ID" envDefault:"11155111"`
	BeaconChainRPCUrl           string `env:"BEACON_CHAIN_RPC_URL" envDefault:"https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30"`
	Blockchain                  string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber             int64  `env:"INIT_BLOCK_NUMBER" envDefault:"4102885"`
	InitBlockHash               string `env:"INIT_BLOCK_HASH" envDefault:"0x9612534dc810c9c51211c77def2db781d7cc7979b0cb076a47c9fc6fb6dc475c"`
	InitBlobBlockNumber         int64  `env:"INIT_BLOB_BLOCK_NUMBER" envDefault:"5687501"`
	InitBlobBlockHash           string `env:"INIT_BLOB_BLOCK_HASH" envDefault:"0x6218666b40fce4153e8f5349ab2f9d2590a601e5a178e4b6d4580094d5c0c2ee"`
	BlobIntervalTime            uint64 `env:"BLOB_INTERVAL_TIME" envDefault:"1800"`
	OutputIntervalTime          uint64 `env:"OUTPUT_INTERVAL_TIME" envDefault:"1800"`
	L2OutputOracleProxyContract string `env:"L2_OUTPUT_ORACLE_PROXY_CONTRACT" envDefault:"0x90E9c4f8a994a250F6aEfd61CAFb4F2e895D458F"`
	BatcherInbox                string `env:"BATCHER_INBOX" envDefault:"0xff00000000000000000000000000000011155420"`
	BatcherSender               string `env:"BATCHER_SENDER" envDefault:"0x8F23BB38F531600e5d8FDDaAEC41F13FaB46E98c"`
	LimitNum                    int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID              uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
	DSType                      string `env:"DSTYPE" envDefault:"arweave"`
	ArweaveWallet               string `env:"B2NODE_ARWEAVE_WALLET" envDefault:"/tmp/wallet/account.json"`
	ArweaveRPC                  string `env:"B2NODE_ARWEAVE_RPC" envDefault:"https://arweave.net"`
	UnisatURL                   string `env:"UNISAT_URL" envDefault:"https://open-api-testnet.unisat.io/"`
	UnisatAuth                  string `env:"UNISAT_PRIVATE_KEY" envDefault:"4cb09e6eec9a0c3ebd07135b3817a3dcedcc9791d50fade4cb564e8ad68a7ac3"`
}

type B2NODEConfig struct {
	ChainID                  int64  `env:"B2NODE_CHAIN_ID" envDefault:"11155111"`
	RPCUrl                   string `env:"B2NODE_RPC_URL" envDefault:"https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30"`
	OpCommitterAddress       string `env:"B2NODE_OP_COMMITTER_ADDRESS" envDefault:"0x270794Fc3ca753CDE033D2AeF9D00EAf71EbC386"`
	OpProposersAddress       string `env:"B2NODE_OP_PROPOSERS_ADDRESS" envDefault:"0xcbC418ce125d806087da0DAb15d6ad50E035a250"`
	OpProposalManagerAddress string `env:"B2NODE_OP_PROPOSAL_MANAGER_ADDRESS" envDefault:"0x837596C1Aa783E3B06C7Efb10a51Fe6699208D1D"`
	Address                  string `env:"B2NODE_CREATOR_ADDRESS" envDefault:"0xb634434CA448c39b05b460dEC51f458EaC1e2759"`
	PrivateKey               string `env:"B2NODE_CREATOR_PRIVATE_KEY" envDefault:"0a81baab0ca0b65d406d68c79945054b092cbe77499ca55c57b3ecfd33f1d551"`
}

type BitcoinRPCConfig struct {
	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"testnet3"`
	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"55968c09fb90a496096bafdeaac0f791f527b17324d1d0e63d3550e68a7b0cc5"`
	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"tb1q6t5py7fqml8patll2jzfc26q7987xqthslyvj4"`
}

var (
	config       *Config
	btcRPCConfig *BitcoinRPCConfig
	b2nodeConfig *B2NODEConfig
)

func GetConfig() *Config {
	if config == nil {
		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		config = cfg
	}
	return config
}

func GetBtcConfig() *BitcoinRPCConfig {
	if btcRPCConfig == nil {
		cfg := &BitcoinRPCConfig{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		btcRPCConfig = cfg
	}
	return btcRPCConfig
}

func GetB2nodeConfig() *B2NODEConfig {
	if b2nodeConfig == nil {
		cfg := &B2NODEConfig{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		b2nodeConfig = cfg
	}
	return b2nodeConfig
}
