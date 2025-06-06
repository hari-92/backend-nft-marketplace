package wallet_converters

import (
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

func ToWalletProtoList(wallets []*walletModels.Wallet) []*pb.Wallet {
	protoWallets := make([]*pb.Wallet, len(wallets))
	for i := range wallets {
		protoWallets[i] = &pb.Wallet{}
	}
	return protoWallets
}
