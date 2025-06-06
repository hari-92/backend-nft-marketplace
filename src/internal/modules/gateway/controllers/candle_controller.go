package gateway_controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	gatewayInstance "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type CandleController struct {
}

func NewCandleController() *CandleController {
	return &CandleController{}
}

// GetCandles: Get candles of a pair with a time frame (GET /candles)
func (c *CandleController) GetCandles(ctx *gin.Context) {
	symbol := ctx.Query("symbol")
	interval := ctx.Query("interval")
	startTime := ctx.Query("start_time")
	endTime := ctx.Query("end_time")

	if symbol == "" || interval == "" || startTime == "" || endTime == "" {
		ctx.JSON(400, gin.H{"error": "Missing required query parameters: symbol, interval, start_time, end_time"})
		return
	}
	startTimeUint64, err := strconv.ParseUint(startTime, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTimeUint64, err := strconv.ParseUint(endTime, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid end_time format"})
		return
	}

	res, err := gatewayInstance.CandleRpcPortGateway.GetCandles(ctx, &pb.GetCandlesRequest{
		Symbol:    symbol,
		Interval:  interval,
		StartTime: startTimeUint64,
		EndTime:   endTimeUint64,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Candles", "data": res})
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
