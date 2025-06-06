package user_scopes

import (
	userFilters "gitlab.com/hari-92/nft-market-server/internal/modules/user/filters"
	"gorm.io/gorm"
)

func UserTokenScope(filter *userFilters.UserTokenFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.UserID != 0 {
			db = db.Where("user_id = ?", filter.UserID)
		}
		if filter.TokenID != 0 {
			db = db.Where("token_id = ?", filter.TokenID)
		}
		if len(filter.InTokenID) != 0 {
			db = db.Where("token_id IN (?)", filter.InTokenID)
		}
		if len(filter.NotInTokenID) != 0 {
			db = db.Where("token_id NOT IN (?)", filter.NotInTokenID)
		}
		if filter.WalletID != 0 {
			db = db.Where("wallet_id = ?", filter.WalletID)
		}
		if len(filter.InWalletID) > 0 {
			db = db.Where("wallet_id IN (?)", filter.InWalletID)
		}
		if len(filter.NotInWalletID) > 0 {
			db = db.Where("wallet_id NOT IN (?)", filter.NotInWalletID)
		}
		return db
	}
}
