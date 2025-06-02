package user_converters

import (
	userResponses "gitlab.com/hari-92/nft-market-server/internal/modules/user/responses"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ResponseToPbConverter struct{}

func NewResponseToPbConverter() *ResponseToPbConverter {
	return &ResponseToPbConverter{}
}

func (c *ResponseToPbConverter) FromGetUserToGetOneUserPb(res *userResponses.GetUserResponse) *pb.GetOneUserResponse {
	if res == nil {
		return nil
	}

	return &pb.GetOneUserResponse{
		Id:    res.ID,
		Email: res.Email,
	}
}

func (c *ResponseToPbConverter) FromCreateUserResponseToCreateUserPb(res *userResponses.CreateUserResponse) *pb.CreateUserResponse {
	if res == nil {
		return nil
	}

	return &pb.CreateUserResponse{
		Id:    res.ID,
		Email: res.Email,
	}
}

func (c *ResponseToPbConverter) FromVerifyUserResponseToLoginPb(res *userResponses.VerifyUserResponse) *pb.LoginResponse {
	if res == nil {
		return nil
	}

	return &pb.LoginResponse{
		Id: res.ID,
	}
}
