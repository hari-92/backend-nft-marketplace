package main

import (
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	"gitlab.com/hari-92/nft-market-server/internal/modules/gateway"
	"gitlab.com/hari-92/nft-market-server/pkg/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.WithDefaultOptions(),
		bootstrap.WithApi(),
		bootstrap.WithGrpcServer(),
		gateway.NewProvider(),
		discovery_client.Provider(),
	).Run()
}
