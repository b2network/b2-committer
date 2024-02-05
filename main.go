package main

import (
	"github.com/b2network/b2committer/internal/handler"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/shopspring/decimal"
)

func main() {
	decimal.DivisionPrecision = 18
	cfg := types.GetConfig()
	btccfg := types.GetBtcConfig()
	b2nodeConfig := types.GetB2nodeConfig()
	log.Init(cfg.LogLevel, cfg.LogFormat)
	log.Infof("config: %v\n", cfg)
	ctx := svc.NewServiceContext(cfg, btccfg, b2nodeConfig)
	handler.Run(ctx)
	log.Info("listener running...\n")
	select {}
}
