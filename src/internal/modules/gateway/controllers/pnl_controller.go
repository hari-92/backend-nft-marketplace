package gateway_controllers

import "github.com/gin-gonic/gin"

type PnlController struct {
}

func NewPnlController() *PnlController {
	return &PnlController{}
}

// GetRealized: Get the realized PnL of a user (GET /pnl/realized)
func (p *PnlController) GetRealized(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl Realized"})
}

// GetUnrealized: Get the unrealized PnL of a user (GET /pnl/unrealized)
func (p *PnlController) GetUnrealized(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl Unrealized"})
}

// GetSummary: Get the summary of the PnL of a user (GET /pnl/summary)
func (p *PnlController) GetSummary(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl Summary"})
}

// GetPortfolio: Get the portfolio of the PnL of a user (GET /pnl/portfolio)
func (p *PnlController) GetPortfolio(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl Portfolio"})
}

// PostRecalculate: Recalculate the PnL of a user (POST /pnl/recalculate)
func (p *PnlController) PostRecalculate(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Recalculate Pnl"})
}

// GetHistory: Get the history of the PnL of a user (GET /pnl/history)
func (p *PnlController) GetHistory(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl History"})
}

// PostValidate: Validate the PnL of a user (POST /pnl/validate)
func (p *PnlController) PostValidate(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Validate Pnl"})
}

// GetPnlPair: Get the PnL of a pair (GET /pnl/pair/:pair_id)
func (p *PnlController) GetPnlPair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pnl Pair"})
}

// GetExport: Export the PnL of a user (GET /pnl/export)
func (p *PnlController) GetExport(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Export Pnl"})
}
