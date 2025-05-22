package pnl_scopes

import (
	pnlFilters "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/filters"
	"gorm.io/gorm"
)

func PnlScope(filter *pnlFilters.PnlFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.UserId != "" {
			db = db.Where("user_id = ?", filter.UserId)
		}
		if filter.PairId != "" {
			db = db.Where("pair_id = ?", filter.PairId)
		}
		if len(filter.InUserIds) > 0 {
			db = db.Where("user_id IN ?", filter.InUserIds)
		}
		if len(filter.InPairIds) > 0 {
			db = db.Where("pair_id IN ?", filter.InPairIds)
		}
		if len(filter.NotInUserIds) > 0 {
			db = db.Where("user_id NOT IN ?", filter.NotInUserIds)
		}
		if len(filter.NotInPairIds) > 0 {
			db = db.Where("pair_id NOT IN ?", filter.NotInPairIds)
		}

		return db
	}
}
