package pnl_services

type PnlService interface {
}

func NewPnlService() PnlService {
	return &pnlService{}
}

type pnlService struct {
}
