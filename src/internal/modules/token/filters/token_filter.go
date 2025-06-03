package token_filters

type TokenFilter struct {
	Id            uint32   `json:"id"`
	Address       string   `json:"address"`
	Symbol        string   `json:"symbol"`
	InSymbols     []string `json:"in_symbols"`
	NotInSymbols  []string `json:"not_in_symbols"`
	ChainID       int      `json:"chain_id"`
	InChainIDs    []int    `json:"in_chain_ids"`
	NotInChainIDs []int    `json:"not_in_chain_ids"`
	IsActive      *bool    `json:"is_active"`
}
