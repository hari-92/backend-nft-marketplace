package pnl

import (
	pnlControllers "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/controllers"
	pnlGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/grpc"
	pnlRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/repositories"
	pnlServices "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(pnlControllers.NewPnlController),

		// Provide services
		fx.Provide(pnlServices.NewPnlService),

		// Provide repositories
		fx.Provide(pnlRepositories.NewPnlRepository),

		// Provide grpc server
		fx.Provide(pnlGrpc.NewGrpcHandler),
		fx.Provide(pnlGrpc.NewGrpcServer),
		fx.Invoke(func(s *pnlGrpc.Server) {}),
	)
}
