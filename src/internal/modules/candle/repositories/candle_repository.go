package candle_respositories

import (
	candleFilters "gitlab.com/hari-92/nft-market-server/internal/modules/candle/filters"
	candleModels "gitlab.com/hari-92/nft-market-server/internal/modules/candle/models"
	candleScopes "gitlab.com/hari-92/nft-market-server/internal/modules/candle/repositories/scopes"
	"gorm.io/gorm"
)

type ICandleRepository interface {
	FindMany(filter *candleFilters.CandlesFilter) (*[]candleModels.Candle, error)
	Create(candle *candleModels.Candle) (*candleModels.Candle, error)
	CreateMany(candles []*candleModels.Candle) error
}

func NewCandleRepository(
	db *gorm.DB,
) ICandleRepository {
	return &CandleRepository{
		db: db,
	}
}

type CandleRepository struct {
	db *gorm.DB
}

func (r *CandleRepository) FindMany(filter *candleFilters.CandlesFilter) (*[]candleModels.Candle, error) {
	var candles []candleModels.Candle
	if err := r.db.Scopes(candleScopes.CandleScope(filter)).Find(&candles).Error; err != nil {
		return nil, err
	}

	return &candles, nil
}

func (r *CandleRepository) Create(candle *candleModels.Candle) (*candleModels.Candle, error) {
	if err := r.db.Create(candle).Error; err != nil {
		return nil, err
	}

	return candle, nil
}

func (r *CandleRepository) CreateMany(candles []*candleModels.Candle) error {
	return r.db.Model(&candleModels.Candle{}).CreateInBatches(candles, len(candles)).Error
}
