package token_requests

type CreateToken struct {
	Address     string
	Symbol      string
	Name        string
	Description string
	Decimals    int
	TotalSupply uint64
	ChainID     uint32
}
