package discovery_client

import (
	"gitlab.com/hari-92/nft-market-server/internal/core/constant"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

var (
	userClient   pb.UserProtoServiceClient
	walletClient pb.WalletProtoServiceClient
	tokenClient  pb.TokenProtoServiceClient
	pairClient   pb.PairProtoServiceClient
	pnlClient    pb.PnlProtoServiceClient
	candleClient pb.CandleProtoServiceClient
)

func GetUserClient() pb.UserProtoServiceClient {
	if userClient != nil {
		return userClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.User)
	if err != nil {
		panic(err)
	}

	userClient = client.(pb.UserProtoServiceClient)
	return userClient
}

func GetWalletClient() pb.WalletProtoServiceClient {
	if walletClient != nil {
		return walletClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.Wallet)
	if err != nil {
		panic(err)
	}

	walletClient = client.(pb.WalletProtoServiceClient)
	return walletClient
}

func GetTokenClient() pb.TokenProtoServiceClient {
	if tokenClient != nil {
		return tokenClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.Token)
	if err != nil {
		panic(err)
	}

	tokenClient = client.(pb.TokenProtoServiceClient)
	return tokenClient
}

func GetPairClient() pb.PairProtoServiceClient {
	if pairClient != nil {
		return pairClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.Pair)
	if err != nil {
		panic(err)
	}

	pairClient = client.(pb.PairProtoServiceClient)
	return pairClient
}

func GetPnlClient() pb.PnlProtoServiceClient {
	if pnlClient != nil {
		return pnlClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.Pnl)
	if err != nil {
		panic(err)
	}

	pnlClient = client.(pb.PnlProtoServiceClient)
	return pnlClient
}

func GetCandleClient() pb.CandleProtoServiceClient {
	if candleClient != nil {
		return candleClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.Candle)
	if err != nil {
		panic(err)
	}

	candleClient = client.(pb.CandleProtoServiceClient)
	return candleClient
}
