package candle_scopes

import (
	candleFilters "gitlab.com/hari-92/nft-market-server/internal/modules/candle/filters"
	"gorm.io/gorm"
)

func CandleScope(filter *candleFilters.CandlesFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.StartTime.IsZero() && filter.EndTime.IsZero() {
			return db
		}

		if !filter.StartTime.IsZero() {
			db = db.Where("time >= ?", filter.StartTime)
		}

		if !filter.EndTime.IsZero() {
			db = db.Where("time <= ?", filter.EndTime)
		}

		if filter.PairID != 0 {
			db = db.Where("pair_id = ?", filter.PairID)
		}

		return db
	}
}
