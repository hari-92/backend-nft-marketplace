package discovery

import (
	"github.com/golibs-starter/golib"
	discoveryGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/discovery/grpc"
	discoveryProperties "gitlab.com/hari-92/nft-market-server/internal/modules/discovery/properties"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		golib.ProvideProps(discoveryProperties.NewServicesDiscoveryProperties),
		fx.Provide(discoveryGrpc.NewGrpcHandler),
		fx.Provide(discoveryGrpc.NewGrpcServer),
		fx.Invoke(func(s *discoveryGrpc.Server) {}),
	)
}
