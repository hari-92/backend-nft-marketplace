package discovery_client

import (
	"github.com/golibs-starter/golib"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/properties"
	"go.uber.org/fx"
)

var (
	discoveryGrpcClient IDiscoveryGrpcClient
)

func Provider() fx.Option {
	return fx.Options(
		golib.ProvideProps(properties.NewGrpcDiscoveryServerProperties),
		fx.Provide(NewDiscoveryClient),
		fx.Invoke(func(client IDiscoveryGrpcClient) {
			discoveryGrpcClient = client
		}),
	)
}
