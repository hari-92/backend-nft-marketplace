package gateway_controllers

import "github.com/gin-gonic/gin"

type CandleController struct {
}

func NewCandleController() *CandleController {
	return &CandleController{}
}

// GetCandles: Get candles of a pair with a time frame (GET /candles/:pair/:time_frame)
func (c *CandleController) GetCandles(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Candles"})
}

// GetCandleStats: Get candle stats of a pair with a time frame (GET /candles/:pair/:time_frame/stats)
func (c *CandleController) GetCandleStats(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Candle Stats"})
}

// GetCandleLatest: Get latest candle of a pair with a time frame (GET /candles/:pair/:time_frame/latest)
func (c *CandleController) GetCandleLatest(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Candle Latest"})
}

// PostCandleValidate: Validate a candle (POST /candles/validate)
func (c *CandleController) PostCandleValidate(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Candle Validate"})
}

// GetCandlesAggregate: Get aggregate candles of a pair (GET /candles/:pair/aggregate)
func (c *CandleController) GetCandlesAggregate(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Candles Aggregate"})
}
