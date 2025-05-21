package gateway_controllers

import (
	"github.com/gin-gonic/gin"
	gatewayInstance "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type PnlController struct {
}

func NewPnlController() *PnlController {
	return &PnlController{}
}

// GetRealized: Get the realized PnL of a user (GET /pnl/realized)
func (p *PnlController) GetRealized(ctx *gin.Context) {
	// TODO: build request
	var getRealizedRequest pb.GetRealizedRequest
	userId := ctx.GetString("user_id")
	getRealizedRequest.UserId = userId
	res, err := gatewayInstance.PnlRpcPortGateway.GetRealized(ctx, &getRealizedRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl Realized", "data": res})
}

// GetUnrealized: Get the unrealized PnL of a user (GET /pnl/unrealized)
func (p *PnlController) GetUnrealized(ctx *gin.Context) {
	// TODO: build request
	getUnrealizedRequest := pb.GetUnrealizedRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&getUnrealizedRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetUnrealized(ctx, &getUnrealizedRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl Unrealized", "data": res})
}

// GetSummary: Get the summary of the PnL of a user (GET /pnl/summary)
func (p *PnlController) GetSummary(ctx *gin.Context) {
	// TODO: build request
	getSummaryRequest := pb.GetSummaryRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&getSummaryRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetSummary(ctx, &getSummaryRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl Summary", "data": res})
}

// GetPortfolio: Get the portfolio of the PnL of a user (GET /pnl/portfolio)
func (p *PnlController) GetPortfolio(ctx *gin.Context) {
	// TODO: build request
	getPortfolioRequest := pb.GetPortfolioRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&getPortfolioRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetPortfolio(ctx, &getPortfolioRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl Portfolio", "data": res})
}

// PostRecalculate: Recalculate the PnL of a user (POST /pnl/recalculate)
func (p *PnlController) PostRecalculate(ctx *gin.Context) {
	// TODO: build request
	postRecalculateRequest := pb.PostRecalculateRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&postRecalculateRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.PostRecalculate(ctx, &postRecalculateRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Recalculate Pnl", "data": res})
}

// GetHistory: Get the history of the PnL of a user (GET /pnl/history)
func (p *PnlController) GetHistory(ctx *gin.Context) {
	// TODO: build request
	getHistoryRequest := pb.GetPnlHistoryRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&getHistoryRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetHistory(ctx, &getHistoryRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl History", "data": res})
}

// PostValidate: Validate the PnL of a user (POST /pnl/validate)
func (p *PnlController) PostValidate(ctx *gin.Context) {
	// TODO: build request
	postValidateRequest := pb.PostValidateRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&postValidateRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.PostValidate(ctx, &postValidateRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Validate Pnl", "data": res})
}

// GetPnlPair: Get the PnL of a pair (GET /pnl/pair/:pair_id)
func (p *PnlController) GetPnlPair(ctx *gin.Context) {
	// TODO: build request
	getPnlPairRequest := pb.GetPnlPairRequest{
		PairId: "1",
	}
	if err := ctx.ShouldBindJSON(&getPnlPairRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetPnlPair(ctx, &getPnlPairRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Pnl Pair", "data": res})
}

// GetExport: Export the PnL of a user (GET /pnl/export)
func (p *PnlController) GetExport(ctx *gin.Context) {
	// TODO: build request
	getExportRequest := pb.GetExportRequest{
		UserId: "1",
	}
	if err := ctx.ShouldBindJSON(&getExportRequest); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	res, err := gatewayInstance.PnlRpcPortGateway.GetExport(ctx, &getExportRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Export Pnl", "data": res})
}
