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
	MySQLDataSource      string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3306)/b2_committer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl               string `env:"RPC_URL" envDefault:"http://haven-b2-nodes.bsquared.network"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"30277"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0x1810ba2a2f66977cc45ad0ef6895393eff479ccfbb854bc8f4aa8154787c1144"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0x67d269191c92Caf3cD7723F116c85e6E9bf55933"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	PrivateKeyHex string `evn:"B2NODE_PRIVATE_KEY_HEX" envDefault:"37927fcde10259a7114a58487cb6303d04c33291ba29bbb8e488eef150e6a59a"`
	Address       string `env:"B2NODE_ADDRESS" envDefault:"ethm1nexknt73vdv6cm3h6ep6u7pe9vg8kr6kqwyl0a"`
	ChainID       string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost      string `env:"B2NODE_GRPC_HOST" envDefault:"127.0.0.1"`
	GRPCPort      uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9090"`
	RPCUrl        string `env:"B2NODE_RPC_URL" envDefault:"http://localhost:8545"`
	CoinDenom     string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton" `
}

type BitcoinRPCConfig struct {
	NetworkName string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
	PrivateKey  string `env:"BITCOIN_PRIVATE_KEY" envDefault:"c545a409ff7f2e66b4bc863a59dcccf0f4387668a92152a058446bcb58a57027"`
	// signet tb1p2rfzw7mdyvkashtls5z6y7e5wlwdfzvjyay9mk2xgsdmzt5zwykq2e0rq8
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
