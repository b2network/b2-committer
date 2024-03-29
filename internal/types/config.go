package types

import (
	"github.com/b2network/b2committer/pkg/log"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat            string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource      string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3306)/b2_committer3?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl               string `env:"RPC_URL" envDefault:"https://habitat-b2-nodes.bsquared.network"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"70000"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0xb2fa3c8011ce25bb1d261403107b58b6aeda8a2af3827e86ad70ee081966d99c"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0xd9571Aaf414b0F51d40D6738813FA4eA782d18B7"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	Address    string `env:"B2NODE_ADDRESS" envDefault:"ethm1nexknt73vdv6cm3h6ep6u7pe9vg8kr6kqwyl0a"`
	ChainID    string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost   string `env:"B2NODE_GRPC_HOST" envDefault:"127.0.0.1"`
	GRPCPort   uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9090"`
	RPCUrl     string `env:"B2NODE_RPC_URL" envDefault:"http://localhost:8545"`
	CoinDenom  string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton"`
	PrivateKey string `env:"B2NODE_PRIVATE_KEY" envDefault:"b2dd35d83b69d0d572616713148e83ba7d7f02fb14f442ffc4246319c61a3fa3"`
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
