package instance

import (
	"sync"

	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/rpc_ports"
)

var (
	UserRpcPortGateway rpc_ports.IUserRpcPorts
	once               sync.Once
)

func NewGatewayInstanceVars() {
	once.Do(func() {
		UserRpcPortGateway = rpc_ports.NewUserRpcPorts()
	})
}
