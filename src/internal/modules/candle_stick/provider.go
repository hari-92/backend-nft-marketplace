package candle_stick

import (
	candleStickControllers "gitlab.com/hari-92/nft-market-server/internal/modules/candle_stick/controllers"
	candleStickGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/candle_stick/grpc"
	candleStickServices "gitlab.com/hari-92/nft-market-server/internal/modules/candle_stick/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(candleStickControllers.NewCandleStickController),
		// Provide services
		fx.Provide(candleStickServices.NewCandleStickService),

		// Provide grpc server
		fx.Provide(candleStickGrpc.NewGrpcServer),
		fx.Provide(candleStickGrpc.NewGrpcHandler),
		fx.Invoke(func(server *candleStickGrpc.Server) {
		}),
	)
}
