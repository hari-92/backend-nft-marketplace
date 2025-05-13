package candle_stick_services

type CandleStickService interface {
}

func NewCandleStickService() CandleStickService {
	return &candleStickService{}
}

type candleStickService struct {
}
