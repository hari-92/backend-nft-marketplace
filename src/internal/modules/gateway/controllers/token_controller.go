package gateway_controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/rpc_ports"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type TokenController struct {
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

// GetTokens Get all and paginated tokens (GET /api/v1/tokens)
func (t *TokenController) GetTokens(ctx *gin.Context) {
	chainId := ctx.Query("chain_id")
	chainIdUint, err := strconv.ParseUint(chainId, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": "Chain ID is not a valid uint32"})
		return
	}
	tokens, err := rpc_ports.NewTokenRpcPorts().GetTokens(ctx, &pb.GetTokensRequest{ChainId: uint32(chainIdUint)})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Get Tokens", "data": tokens})
}

// GetToken Get a token by id (GET /api/v1/tokens/:id)
func (t *TokenController) GetToken(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": "Token ID is not a valid uint32"})
		return
	}
	token, err := rpc_ports.NewTokenRpcPorts().GetToken(ctx, &pb.GetTokenRequest{Id: uint32(idUint)})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Get Token", "data": token})
}

// PostToken Create a new token (POST /api/v1/tokens)
func (t *TokenController) PostToken(ctx *gin.Context) {
	var req pb.PostTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	res, err := rpc_ports.NewTokenRpcPorts().PostToken(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post Token", "data": res})
}

// PutToken Update token info (PUT /api/v1/tokens/:id)
func (t *TokenController) PutToken(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": "Token ID is not a valid uint32"})
		return
	}
	var req pb.PutTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	req.Id = uint32(idUint)
	res, err := rpc_ports.NewTokenRpcPorts().PutToken(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Put Token", "data": res})
}

// DeleteToken Delete a token (DELETE /api/v1/tokens/:id)
func (t *TokenController) DeleteToken(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": "Token ID is not a valid uint32"})
		return
	}
	req := pb.DeleteTokenRequest{
		Id: uint32(idUint),
	}
	res, err := rpc_ports.NewTokenRpcPorts().DeleteToken(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Delete Token", "data": res})
}

// PostValidateToken Validate a token (POST /api/v1/tokens/validate)
func (t *TokenController) PostValidateToken(ctx *gin.Context) {
	var req pb.PostValidateTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	res, err := rpc_ports.NewTokenRpcPorts().PostValidateToken(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post Validate Token", "data": res})
}
