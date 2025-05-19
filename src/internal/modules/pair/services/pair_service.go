package pair_services

type PairService interface {
}

func NewPairService() PairService {
	return &pairService{}
}

type pairService struct {
}
