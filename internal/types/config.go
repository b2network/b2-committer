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
	MySQLDataSource            string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3366)/b2_committer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns          int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns          int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime       int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl                     string `env:"RPC_URL" envDefault:"https://sepolia.drpc.org"`
	Blockchain                 string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber            int64  `env:"INIT_BLOCK_NUMBER" envDefault:"5158900"`
	InitBlockHash              string `env:"INIT_BLOCK_HASH" envDefault:"0xc05e1c7dd54b60c5bc13ad09bad119f49a2fc82cc60215dd7eb9b1d68e147ab0"`
	PolygonSequenceContract    string `env:"POLYGON_SEQUENCE_CONTRACT" envDefault:"0xA13Ddb14437A8F34897131367ad3ca78416d6bCa"`
	PolygonVerifyBatchContract string `env:"POLYGON_VERIFY_BATCH_CONTRACT" envDefault:"0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff"`
	LimitNum                   int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID             uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	ChainID          int64  `env:"B2NODE_CHAIN_ID" envDefault:"11155111"`
	RPCUrl           string `env:"B2NODE_RPC_URL" envDefault:"https://eth-sepolia.g.alchemy.com/v2/lV2e-64nNnEMUA7UG0IT0uwjzlxEI512"`
	CommitterAddress string `env:"B2NODE_COMMITTER_ADDRESS" envDefault:"0x12BBD3f7EF1ABEd6B9DB12A3dE77b00aE10618E0"`
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
