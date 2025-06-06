package candle_services

import (
	common_producers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	commonProducersEventCandle "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/candle"
	candleModels "gitlab.com/hari-92/nft-market-server/internal/modules/candle/models"
	candleRepository "gitlab.com/hari-92/nft-market-server/internal/modules/candle/repositories"
	candleRequests "gitlab.com/hari-92/nft-market-server/internal/modules/candle/requests"
	candleResponses "gitlab.com/hari-92/nft-market-server/internal/modules/candle/responses"
)

type ICandleService interface {
	GetCandles(req *candleRequests.GetCandlesRequest) (*[]candleResponses.GetCandlesResponse, error)
	Create(req *candleRequests.CreateCandleRequest) (*candleResponses.CreateCandleResponse, error)
}

func NewCandleService(
	candleRepository candleRepository.ICandleRepository,
) ICandleService {
	return &CandleService{
		candleRepository: candleRepository,
	}
}

type CandleService struct {
	candleRepository candleRepository.ICandleRepository
}

func (s *CandleService) GetCandles(req *candleRequests.GetCandlesRequest) (*[]candleResponses.GetCandlesResponse, error) {
	return nil, nil
}

func (s *CandleService) Create(req *candleRequests.CreateCandleRequest) (*candleResponses.CreateCandleResponse, error) {
	res, err := s.candleRepository.Create(&candleModels.Candle{
		PairID: req.PairID,
	})
	if err != nil {
		return nil, err
	}

	common_producers.GetProducerHubInstance().CreatedCandleEvent(&commonProducersEventCandle.CreatedCandleEvent{
		PairID: res.PairID,
	})

	return &candleResponses.CreateCandleResponse{
		PairID: res.PairID,
	}, nil
}
