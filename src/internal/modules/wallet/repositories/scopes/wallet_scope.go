package wallet_scopes

import (
	walletFilters "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/filters"
	"gorm.io/gorm"
)

func WalletScope(filter *walletFilters.WalletFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.UserID != 0 {
			db = db.Where("user_id = ?", filter.UserID)
		}

		if len(filter.InUserID) > 0 {
			db = db.Where("user_id IN ?", filter.InUserID)
		}

		if len(filter.NotInUserID) > 0 {
			db = db.Where("user_id NOT IN ?", filter.NotInUserID)
		}

		if filter.ChainID != 0 {
			db = db.Where("chain_id = ?", filter.ChainID)
		}

		if len(filter.InChainID) > 0 {
			db = db.Where("chain_id IN ?", filter.InChainID)
		}

		if len(filter.NotInChainID) > 0 {
			db = db.Where("chain_id NOT IN ?", filter.NotInChainID)
		}

		if filter.Address != "" {
			db = db.Where("address = ?", filter.Address)
		}

		if filter.TokenID != 0 {
			db = db.Where("token_id = ?", filter.TokenID)
		}

		if len(filter.InTokenID) > 0 {
			db = db.Where("token_id IN ?", filter.InTokenID)
		}
		
		if len(filter.NotInTokenID) > 0 {
			db = db.Where("token_id NOT IN ?", filter.NotInTokenID)
		}

		return db
	}
}
