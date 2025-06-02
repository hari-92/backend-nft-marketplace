package common_producers_event

type InitTokenEvent struct {
	ID          uint32 `json:"id"`
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	TokenName   string `json:"name"`
	Description string `json:"description"`
	TotalSupply uint64 `json:"total_supply"`
	ChainID     uint32 `json:"chain_id"`
}

func (i InitTokenEvent) Name() string {
	return "InitTokenEvent"
}

func (i InitTokenEvent) Data() interface{} {
	return i
}
