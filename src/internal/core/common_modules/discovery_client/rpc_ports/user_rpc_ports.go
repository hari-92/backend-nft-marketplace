package rpc_ports

type IUserRpcPorts interface {
}

type userRpcPorts struct {
}

func NewUserRpcPorts() IUserRpcPorts {
	return &userRpcPorts{}
}
