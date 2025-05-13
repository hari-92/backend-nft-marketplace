package pnl_controllers

import (
	pnlServices "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/services"
)

type PnlController struct {
	pnlService pnlServices.PnlService
}

func NewPnlController(pnlService pnlServices.PnlService) *PnlController {
	return &PnlController{pnlService: pnlService}
}
