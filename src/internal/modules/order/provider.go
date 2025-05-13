package order

import (
	orderControllers "gitlab.com/hari-92/nft-market-server/internal/modules/order/controllers"
	orderGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/order/grpc"
	orderServices "gitlab.com/hari-92/nft-market-server/internal/modules/order/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(orderControllers.NewOrderController),

		// Provide services
		fx.Provide(orderServices.NewOrderService),

		// Provide grpc server
		fx.Provide(orderGrpc.NewGrpcHandler),
		fx.Provide(orderGrpc.NewGrpcServer),
		fx.Invoke(func(server *orderGrpc.Server) {
		}),
	)
}
