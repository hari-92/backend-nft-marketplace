package user_converters

import (
	userRequests "gitlab.com/hari-92/nft-market-server/internal/modules/user/requests"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type PbToRequestConverter struct{}

func NewPbToRequestConverter() *PbToRequestConverter {
	return &PbToRequestConverter{}
}

func (c *PbToRequestConverter) FromGetOneUserToGetUserRequest(proto *pb.GetOneUserRequest) *userRequests.GetUserRequest {
	if proto == nil {
		return nil
	}

	return &userRequests.GetUserRequest{
		ID:    proto.Id,
		Email: proto.Email,
	}
}

func (c *PbToRequestConverter) FromIsExistUserToGetUserRequest(proto *pb.IsExistUserRequest) *userRequests.IsExistUserRequest {
	if proto == nil {
		return nil
	}

	return &userRequests.IsExistUserRequest{
		Username: proto.Username,
	}
}

func (c *PbToRequestConverter) FromCreateUserToCreateUserRequest(proto *pb.CreateUserRequest) *userRequests.CreateUserRequest {
	if proto == nil {
		return nil
	}

	return &userRequests.CreateUserRequest{
		Email:    proto.Email,
		Password: proto.Password,
	}
}
