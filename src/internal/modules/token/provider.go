package token

import (
	tokenControllers "gitlab.com/hari-92/nft-market-server/internal/modules/token/controllers"
	tokenGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/token/grpc"
	tokenServices "gitlab.com/hari-92/nft-market-server/internal/modules/token/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		fx.Provide(
			// Provide controllers
			fx.Provide(tokenControllers.NewTokenController),

			// Provide services
			fx.Provide(tokenServices.NewTokenService),

			// Provide grpc server
			fx.Provide(tokenGrpc.NewGrpcServer),
			fx.Provide(tokenGrpc.NewHandler),
		),
	)
}
