package token_grpc

import pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"

type Handler struct {
	pb.UnimplementedTokenProtoServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}
