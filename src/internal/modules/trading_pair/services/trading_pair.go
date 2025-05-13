package trading_pair_services

type TradingPairService interface {
}

func NewTradingPairService() TradingPairService {
	return &tradingPairService{}
}

type tradingPairService struct {
}
