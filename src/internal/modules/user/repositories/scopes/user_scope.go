package user_scopes

import (
	userFilters "gitlab.com/hari-92/nft-market-server/internal/modules/user/filters"
	"gorm.io/gorm"
)

func UserScope(filter *userFilters.UserFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db = db.Where("id = ?", filter.ID)
		}

		if filter.Email != "" {
			db = db.Where("email = ?", filter.Email)
		}

		if len(filter.InIDs) > 0 {
			db = db.Where("id IN ?", filter.InIDs)
		}

		if len(filter.NotInIDs) > 0 {
			db = db.Where("id NOT IN ?", filter.NotInIDs)
		}

		return db
	}
}
