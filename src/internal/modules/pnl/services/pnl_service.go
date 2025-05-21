package pnl_services

import (
	pnlRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/repositories"
)

type PnlService interface {
}

func NewPnlService(pnlRepository pnlRepositories.PnlRepository) PnlService {
	return &pnlService{pnlRepository: pnlRepository}
}

type pnlService struct {
	pnlRepository pnlRepositories.PnlRepository
}
