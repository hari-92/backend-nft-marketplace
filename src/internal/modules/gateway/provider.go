package gateway

import (
	"gitlab.com/hari-92/nft-market-server/internal/modules/gateway/routers"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		fx.Invoke(routers.RegisterHandler),
	)
}
