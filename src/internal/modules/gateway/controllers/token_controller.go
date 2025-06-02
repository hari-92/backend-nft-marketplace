package gateway_controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/rpc_ports"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type TokenController struct {
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

// GetTokens Get all and paginated tokens (GET /tokens)
func (t *TokenController) GetTokens(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Tokens"})
}

// GetToken Get a token by id (GET /tokens/:token_id)
func (t *TokenController) GetToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Token"})
}

// PostToken Create a new token (POST /tokens)
func (t *TokenController) PostToken(ctx *gin.Context) {
	res, err := rpc_ports.NewTokenRpcPorts().PostToken(ctx, &pb.PostTokenRequest{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post Token", "data": res})
}

// PutToken Update token info (PUT /tokens/:token_id)
func (t *TokenController) PutToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Put Token"})
}

// DeleteToken Delete a token (DELETE /tokens/:token_id)
func (t *TokenController) DeleteToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Delete Token"})
}

// PostValidateToken Validate a token (POST /tokens/validate)
func (t *TokenController) PostValidateToken(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Validate Token"})
}
