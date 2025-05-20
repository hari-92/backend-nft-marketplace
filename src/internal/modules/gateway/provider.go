package gateway

import (
	gatewayControllers "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/controllers"
	gatewayGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/grpc"
	gatewayInstance "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
	gatewayRouters "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/routers"
	gatewayServices "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(gatewayControllers.NewUserController),
		fx.Provide(gatewayControllers.NewAuthController),
		fx.Provide(gatewayControllers.NewWalletController),
		fx.Provide(gatewayControllers.NewTokenController),
		fx.Provide(gatewayControllers.NewPairController),
		fx.Provide(gatewayControllers.NewPnlController),
		fx.Provide(gatewayControllers.NewOrderController),
		fx.Provide(gatewayControllers.NewCandleController),

		// Provide service
		fx.Provide(gatewayServices.NewAuthService),

		// Provide router handler
		fx.Invoke(gatewayRouters.RegisterRoutes),
		fx.Invoke(gatewayRouters.RegisterHandler),

		fx.Provide(gatewayGrpc.NewGrpcHandler),
		fx.Provide(gatewayGrpc.NewGrpcServer),
		fx.Invoke(func(s *gatewayGrpc.Server) {}),

		fx.Invoke(gatewayInstance.NewGatewayInstanceVars),
	)
}
