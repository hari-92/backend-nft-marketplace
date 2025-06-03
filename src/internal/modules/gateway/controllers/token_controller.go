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
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Tokens"})
}

// GetToken Get a token by id (GET /api/v1/tokens/:token_id)
func (t *TokenController) GetToken(ctx *gin.Context) {
	tokenID := ctx.Param("token_id")
	tokenIDUint, err := strconv.ParseUint(tokenID, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": "Token ID is not a valid uint32"})
		return
	}
	token, err := rpc_ports.NewTokenRpcPorts().GetToken(ctx, &pb.GetTokenRequest{TokenId: uint32(tokenIDUint)})
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

// PutToken Update token info (PUT /api/v1/tokens/:token_id)
func (t *TokenController) PutToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Put Token"})
}

// DeleteToken Delete a token (DELETE /api/v1/tokens/:token_id)
func (t *TokenController) DeleteToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Delete Token"})
}

// PostValidateToken Validate a token (POST /api/v1/tokens/validate)
func (t *TokenController) PostValidateToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Validate Token"})
}
