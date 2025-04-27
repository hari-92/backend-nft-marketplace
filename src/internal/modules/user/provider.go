package user

import (
	"gitlab.com/hari-92/nft-market-server/internal/modules/user/routers"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		fx.Invoke(routers.RegisterHandler),
		fx.Invoke(routers.RegisterRoutes),
	)
}
