package token_scopes

import (
	tokenFilters "gitlab.com/hari-92/nft-market-server/internal/modules/token/filters"
	"gorm.io/gorm"
)

func TokenScope(filter *tokenFilters.TokenFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.Id != 0 {
			db = db.Where("id = ?", filter.Id)
		}
		if filter.Symbol != "" {
			db = db.Where("symbol = ?", filter.Symbol)
		}
		if filter.Address != "" {
			db = db.Where("address = ?", filter.Address)
		}
		if filter.ChainID != 0 {
			db = db.Where("chain_id = ?", filter.ChainID)
		}
		if len(filter.InSymbols) > 0 {
			db = db.Where("symbol IN ?", filter.InSymbols)
		}
		if len(filter.NotInSymbols) > 0 {
			db = db.Where("symbol NOT IN ?", filter.NotInSymbols)
		}
		if len(filter.InChainIDs) > 0 {
			db = db.Where("chain_id IN ?", filter.InChainIDs)
		}
		if len(filter.NotInChainIDs) > 0 {
			db = db.Where("chain_id NOT IN ?", filter.NotInChainIDs)
		}
		if filter.IsActive != nil {
			db = db.Where("is_active = ?", filter.IsActive)
		}

		return db
	}
}
