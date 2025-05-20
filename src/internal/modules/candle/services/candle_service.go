package candle_services

type CandleService interface {
}

func NewCandleService() CandleService {
	return &candleService{}
}

type candleService struct {
}
