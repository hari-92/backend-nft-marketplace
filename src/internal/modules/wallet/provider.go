package wallet

import (
	walletControllers "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/controllers"
	walletGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/grpc"
	walletRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
	walletServices "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide controllers
		fx.Provide(walletControllers.NewWalletController),

		// Provide services
		fx.Provide(walletServices.NewWalletService),
		fx.Provide(walletServices.NewUserService),

		// Provide repositories
		fx.Provide(walletRepositories.NewWalletRepository),

		//Provide grpc server
		fx.Provide(walletGrpc.NewGrpcHandler),
		fx.Provide(walletGrpc.NewGrpcServer),
		fx.Invoke(func(s *walletGrpc.Server) {}),
	)
}
