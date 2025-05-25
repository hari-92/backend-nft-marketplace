package user_converters

import (
	userResponses "gitlab.com/hari-92/nft-market-server/internal/modules/user/responses"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ResponseToPbConverter struct{}

func NewResponseToPbConverter() *ResponseToPbConverter {
	return &ResponseToPbConverter{}
}

func (c *ResponseToPbConverter) FromGetUserToGetOneUserPb(proto *userResponses.GetUserResponse) *pb.GetOneUserResponse {
	if proto == nil {
		return nil
	}

	return &pb.GetOneUserResponse{
		Id:    proto.ID,
		Email: proto.Email,
	}
}
