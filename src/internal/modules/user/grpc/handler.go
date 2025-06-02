package user_grpc

import (
	"context"

	userConverters "gitlab.com/hari-92/nft-market-server/internal/modules/user/converters"
	userServices "gitlab.com/hari-92/nft-market-server/internal/modules/user/services"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedUserProtoServiceServer
	pbToRequestConverter  *userConverters.PbToRequestConverter
	responseToPbConverter *userConverters.ResponseToPbConverter
	userService           userServices.UserService
}

func NewGrpcHandler(
	userService userServices.UserService,
) *Handler {
	return &Handler{
		pbToRequestConverter:  userConverters.NewPbToRequestConverter(),
		responseToPbConverter: userConverters.NewResponseToPbConverter(),
		userService:           userService,
	}
}

func (h *Handler) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hello world from user service",
	}, nil
}

func (h *Handler) IsExistUser(ctx context.Context, req *pb.IsExistUserRequest) (*pb.IsExistUserResponse, error) {
	res, err := h.userService.IsExist(h.pbToRequestConverter.FromIsExistUserToGetUserRequest(req))
	if err != nil {
		return nil, err
	}
	return &pb.IsExistUserResponse{
		IsExist: res,
	}, nil
}

func (h *Handler) GetOneUser(ctx context.Context, req *pb.GetOneUserRequest) (*pb.GetOneUserResponse, error) {
	res, err := h.userService.GetUser(h.pbToRequestConverter.FromGetOneUserToGetUserRequest(req))
	if err != nil {
		return nil, err
	}
	return h.responseToPbConverter.FromGetUserToGetOneUserPb(res), nil
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := h.userService.CreateUser(h.pbToRequestConverter.FromCreateUserToCreateUserRequest(req))
	if err != nil {
		return nil, err
	}
	return h.responseToPbConverter.FromCreateUserResponseToCreateUserPb(res), nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := h.userService.Verify(h.pbToRequestConverter.FromLoginToVerifyUserRequest(req))
	if err != nil {
		return nil, err
	}
	return h.responseToPbConverter.FromVerifyUserResponseToLoginPb(res), nil
}
