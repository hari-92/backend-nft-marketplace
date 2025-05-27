package user

import (
	userGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/user/grpc"
	userRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/user/repositories"
	userRouters "gitlab.com/hari-92/nft-market-server/internal/modules/user/routers"
	userServices "gitlab.com/hari-92/nft-market-server/internal/modules/user/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers

		// Provide services
		fx.Provide(userServices.NewUserService),
		// Provide repositories
		fx.Provide(userRepositories.NewUserRepository),

		// Provide grpc handlers
		fx.Provide(userGrpc.NewGrpcHandler),
		fx.Provide(userGrpc.NewGrpcServer),

		fx.Invoke(userRouters.RegisterHandler),
		fx.Invoke(userRouters.RegisterRoutes),
		fx.Invoke(func(s *userGrpc.Server) {}),
	)
}
