package main

import (
	"gitlab.com/hari-92/nft-market-server/internal/modules/pair"
	"gitlab.com/hari-92/nft-market-server/pkg/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.WithDefaultOptions(),
		bootstrap.WithApi(),
		bootstrap.WithGrpcServer(),
		pair.NewProvider(),
	).Run()
}
