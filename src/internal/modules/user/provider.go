package user

import (
	userGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/user/grpc"
	userRouters "gitlab.com/hari-92/nft-market-server/internal/modules/user/routers"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		fx.Invoke(userRouters.RegisterHandler),
		fx.Invoke(userRouters.RegisterRoutes),

		fx.Provide(userGrpc.NewGrpcHandler),
		fx.Provide(userGrpc.NewGrpcServer),
		fx.Invoke(func(s *userGrpc.Server) {}),
	)
}
