package gateway_instance

import (
	"sync"

	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/rpc_ports"
)

var (
	once                 sync.Once
	UserRpcPortGateway   rpc_ports.IUserRpcPorts
	WalletRpcPortGateway rpc_ports.IWalletRpcPorts
	TokenRpcPortGateway  rpc_ports.ITokenRpcPorts
	PairRpcPortGateway   rpc_ports.IPairRpcPorts
	PnlRpcPortGateway    rpc_ports.IPnlRpcPorts
)

func NewGatewayInstanceVars() {
	once.Do(func() {
		UserRpcPortGateway = rpc_ports.NewUserRpcPorts()
		WalletRpcPortGateway = rpc_ports.NewWalletRpcPorts()
		TokenRpcPortGateway = rpc_ports.NewTokenRpcPorts()
		PairRpcPortGateway = rpc_ports.NewPairRpcPorts()
		PnlRpcPortGateway = rpc_ports.NewPnlRpcPorts()
	})
}
