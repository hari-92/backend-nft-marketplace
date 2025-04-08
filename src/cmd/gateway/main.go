package main

import (
	"gitlab.com/hari-92/nft-market-server/internal/modules/gateway"
	"gitlab.com/hari-92/nft-market-server/pkg/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.WithDefaultOptions(),
		bootstrap.WithApi(),

		gateway.NewProvider(),
	).Run()
}
