package pair

import (
	pairControllers "gitlab.com/hari-92/nft-market-server/internal/modules/pair/controllers"
	pairGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/pair/grpc"
	pairServices "gitlab.com/hari-92/nft-market-server/internal/modules/pair/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(pairControllers.NewPairController),

		// Provide services
		fx.Provide(pairServices.NewPairService),

		// Provide grpc server
		fx.Provide(pairGrpc.NewGrpcHandler),
		fx.Provide(pairGrpc.NewGrpcServer),
		fx.Invoke(func(s *pairGrpc.Server) {}),
	)
}
