package wallet

import (
	"github.com/golibs-starter/golib"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	commonProducers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	core_properties "gitlab.com/hari-92/nft-market-server/internal/core/properties"
	walletConsumers "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/consumers"
	walletControllers "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/controllers"
	walletGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/grpc"
	walletRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
	walletServices "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		// Provide properties
		golib.ProvideProps(core_properties.NewSecretKeyProperties),

		// Provide controllers
		fx.Provide(walletControllers.NewWalletController),

		// Provide services
		fx.Provide(walletServices.NewWalletService),
		fx.Provide(walletServices.NewUserService),

		// Provide repositories
		fx.Provide(walletRepositories.NewWalletRepository),

		// Provide consumers
		golibmsg.ProvideConsumer(walletConsumers.NewUserCreatedConsumer),

		// Provide common producers
		fx.Provide(commonProducers.Provider),

		//Provide grpc server
		fx.Provide(walletGrpc.NewGrpcHandler),
		fx.Provide(walletGrpc.NewGrpcServer),
		fx.Invoke(func(s *walletGrpc.Server) {}),
	)
}
