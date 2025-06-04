package token_filters

type TokenFilter struct {
	Id            uint32   `json:"id"`
	Address       string   `json:"address"`
	Symbol        string   `json:"symbol"`
	InSymbols     []string `json:"in_symbols"`
	NotInSymbols  []string `json:"not_in_symbols"`
	ChainID       uint32   `json:"chain_id"`
	InChainIDs    []uint32 `json:"in_chain_ids"`
	NotInChainIDs []uint32 `json:"not_in_chain_ids"`
	IsActive      *bool    `json:"is_active"`
}
