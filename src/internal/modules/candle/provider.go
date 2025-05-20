package candle_stick

import (
	candleControllers "gitlab.com/hari-92/nft-market-server/internal/modules/candle/controllers"
	candleGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/candle/grpc"
	candleServices "gitlab.com/hari-92/nft-market-server/internal/modules/candle/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(candleControllers.NewCandleController),
		// Provide services
		fx.Provide(candleServices.NewCandleService),

		// Provide grpc server
		fx.Provide(candleGrpc.NewGrpcServer),
		fx.Provide(candleGrpc.NewGrpcHandler),
		fx.Invoke(func(server *candleGrpc.Server) {
		}),
	)
}
