package candle_grpc

import (
	"context"

	candleRequests "gitlab.com/hari-92/nft-market-server/internal/modules/candle/requests"
	candleServices "gitlab.com/hari-92/nft-market-server/internal/modules/candle/services"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedCandleProtoServiceServer
	candleService candleServices.ICandleService
}

func NewGrpcHandler(
	candleService candleServices.ICandleService,
) *Handler {
	return &Handler{
		candleService: candleService,
	}
}

func (h *Handler) GetCandles(ctx context.Context, req *pb.GetCandlesRequest) (*pb.GetCandlesResponse, error) {
	res, err := h.candleService.GetCandles(&candleRequests.GetCandlesRequest{
		Symbol:    req.Symbol,
		Interval:  req.Interval,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})

	if err != nil {
		return nil, err
	}
	if res == nil {
		return &pb.GetCandlesResponse{}, nil
	}

	var candles []*pb.Candle
	for _, candle := range *res {
		candles = append(candles, &pb.Candle{
			OpenTime:         candle.OpenTime,
			CloseTime:        candle.CloseTime,
			Open:             candle.Open,
			Close:            candle.Close,
			High:             candle.High,
			Low:              candle.Low,
			Volume:           candle.Volume,
			QuoteAssetVolume: candle.QuoteAssetVolume,
			NumberOfTrades:   candle.NumberOfTrades,
		})
	}
	if len(candles) == 0 {
		return &pb.GetCandlesResponse{}, nil
	}
	return &pb.GetCandlesResponse{
		Candles: candles,
	}, nil
}
