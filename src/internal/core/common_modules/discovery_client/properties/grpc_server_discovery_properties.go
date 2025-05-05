package properties

import "github.com/golibs-starter/golib/config"

type GrpcServerDiscoveryProperties struct {
	Host string `json:"host"`
	Port uint64 `json:"port"`
}

func (p GrpcServerDiscoveryProperties) Prefix() string {
	return "app.grpcServerDiscovery"
}

func NewGrpcDiscoveryServerProperties(loader config.Loader) (*GrpcServerDiscoveryProperties, error) {
	props := GrpcServerDiscoveryProperties{}
	err := loader.Bind(&props)
	return &props, err
}
