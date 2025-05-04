package discovery

import (
	"github.com/golibs-starter/golib"
	"gitlab.com/hari-92/nft-market-server/internal/modules/discovery/grpc"
	"gitlab.com/hari-92/nft-market-server/internal/modules/discovery/properties"
	"go.uber.org/fx"
)

func NewProvider() fx.Option {
	return fx.Options(
		golib.ProvideProps(properties.NewServicesDiscoveryProperties),
		fx.Provide(grpc.NewGrpcHandler),
		fx.Provide(grpc.NewGrpcServer),
		fx.Invoke(func(s *grpc.Server) {}),
	)
}
