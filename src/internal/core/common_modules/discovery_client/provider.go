package discovery_client

import (
	"github.com/golibs-starter/golib"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/properties"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/rpc_ports"
	"go.uber.org/fx"
)

var (
	discoveryGrpcClient IDiscoveryGrpcClient
	UserRpcPorts        rpc_ports.IUserRpcPorts
)

func Provider() fx.Option {
	return fx.Options(
		golib.ProvideProps(properties.NewGrpcDiscoveryServerProperties),
		fx.Provide(NewDiscoveryClient),
		fx.Invoke(func(client IDiscoveryGrpcClient) {
			discoveryGrpcClient = client
		}),
		fx.Provide(rpc_ports.NewUserRpcPorts),
		fx.Invoke(func(ports rpc_ports.IUserRpcPorts) {
			UserRpcPorts = ports
		}),
	)
}
