package rpc_ports

import (
	"context"

	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type IWalletRpcPorts interface {
	GetBalance(ctx context.Context, userID uint32) ([]*pb.Wallet, error)
}

type walletRpcPorts struct {
}

func NewWalletRpcPorts() IWalletRpcPorts {
	return &walletRpcPorts{}
}

func (w *walletRpcPorts) GetBalance(ctx context.Context, userID uint32) ([]*pb.Wallet, error) {
	res, err := discovery_client.GetWalletClient().GetBalance(ctx, &pb.GetBalanceRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return res.GetWallets(), nil
}
