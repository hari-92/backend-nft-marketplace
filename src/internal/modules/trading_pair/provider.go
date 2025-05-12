package trading_pair

import (
	tradingPairControllers "gitlab.com/hari-92/nft-market-server/internal/modules/trading_pair/controllers"
	tradingPairGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/trading_pair/grpc"
	tradingPairServices "gitlab.com/hari-92/nft-market-server/internal/modules/trading_pair/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(tradingPairControllers.NewTradingPairController),

		// Provide services
		fx.Provide(tradingPairServices.NewTradingPairService),

		// Provide grpc server
		fx.Provide(tradingPairGrpc.NewGrpcHandler),
		fx.Provide(tradingPairGrpc.NewGrpcServer),
		fx.Invoke(func(s *tradingPairGrpc.Server) {}),
	)
}
