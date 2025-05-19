package pair_controllers

import (
	pairServices "gitlab.com/hari-92/nft-market-server/internal/modules/pair/services"
)

type PairController struct {
	pairService pairServices.PairService
}

func NewPairController(pairService pairServices.PairService) *PairController {
	return &PairController{
		pairService: pairService,
	}
}
