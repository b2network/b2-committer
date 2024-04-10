package types

import (
	"github.com/b2network/b2committer/pkg/log"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat                  string `env:"LOG_FORMAT" envDefault:"console"`
	PostgresqlDataSource       string `env:"POSTGRESQL_DATA_SOURCE" envDefault:"host=localhost port=5433 user=postgres password=postgres dbname=b2_committer sslmode=disable"`
	MySQLDataSource            string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3366)/b2_committer2?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns          int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns          int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime       int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl                     string `env:"RPC_URL" envDefault:"https://habitat-hub-rpc.bsquared.network"`
	Blockchain                 string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber            int64  `env:"INIT_BLOCK_NUMBER" envDefault:"0"`
	InitBlockHash              string `env:"INIT_BLOCK_HASH" envDefault:"0x49a34ca5e11dbcd56b4cad717a98efe5c76552a2dcfbdcd693a248e2fd5f8d51"`
	PolygonSequenceContract    string `env:"POLYGON_SEQUENCE_CONTRACT" envDefault:"0xa6AAdA6845b2083ff6812bAc773038442e7f4dE6"`
	PolygonVerifyBatchContract string `env:"POLYGON_VERIFY_BATCH_CONTRACT" envDefault:"0xDdee8ddfA81F5E36373637240038DCCC14529BF7"`
	LimitNum                   int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID             uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
	BatchPath                  string `env:"BATCH_PATH" envDefault:"batchFile"`
	ArweaveWallet              string `env:"B2NODE_ARWEAVE_WALLET" envDefault:"./wallet/account.json"`
	ArweaveRPC                 string `env:"B2NODE_ARWEAVE_RPC" envDefault:"https://arweave.net"`
}

type B2NODEConfig struct {
	ChainID          int64  `env:"B2NODE_CHAIN_ID" envDefault:"1113"`
	RPCUrl           string `env:"B2NODE_RPC_URL" envDefault:"https://habitat-hub-rpc.bsquared.network"`
	CommitterAddress string `env:"B2NODE_COMMITTER_ADDRESS" envDefault:"0x85D40bDc724bcabF6D17d8343a74e0d916dfD40D"`
	Address          string `env:"B2NODE_CREATOR_ADDRESS" envDefault:"0xb634434CA448c39b05b460dEC51f458EaC1e2759"`
	PrivateKey       string `env:"B2NODE_CREATOR_PRIVATE_KEY" envDefault:"0a81baab0ca0b65d406d68c79945054b092cbe77499ca55c57b3ecfd33f1d551"`
}

type BitcoinRPCConfig struct {
	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"c545a409ff7f2e66b4bc863a59dcccf0f4387668a92152a058446bcb58a57027"`
	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"tb1pvhr4e58yatk9uve22rr5umxs0jh9g0j0gtcj0ry2wf23lddhjptsf6c360"`
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
