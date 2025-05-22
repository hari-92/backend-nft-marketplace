package token

import (
	tokenControllers "gitlab.com/hari-92/nft-market-server/internal/modules/token/controllers"
	tokenGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/token/grpc"
	tokenRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories"
	tokenServices "gitlab.com/hari-92/nft-market-server/internal/modules/token/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(tokenControllers.NewTokenController),

		// Provide services
		fx.Provide(tokenServices.NewTokenService),

		// Provide repositories
		fx.Provide(tokenRepositories.NewTokenRepository),

		// Provide grpc server
		fx.Provide(tokenGrpc.NewGrpcServer),
		fx.Provide(tokenGrpc.NewHandler),
	)
}
