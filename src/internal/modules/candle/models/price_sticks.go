package candle_models

import "time"

type PriceSticks struct {
	Time   time.Time `json:"time" gorm:"type:timestamptz;not null;primaryKey"`
	PairID uint32    `json:"pair_id" gorm:"type:bigint;not null;index"`
	Price  float64   `json:"price" gorm:"type:decimal(18,8);not null"`
	Volume float64   `json:"volume" gorm:"type:decimal(18,8);not null"`
}
