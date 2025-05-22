package candle_models

import (
	"time"
)

type Candle struct {
	Time      time.Time `json:"time" gorm:"type:timestamptz;not null;primaryKey"`
	PairID    uint64    `json:"pair_id" gorm:"type:bigint;not null;index"`
	Open      float64   `json:"open" gorm:"type:decimal(18,8);not null"`
	High      float64   `json:"high" gorm:"type:decimal(18,8);not null"`
	Low       float64   `json:"low" gorm:"type:decimal(18,8);not null"`
	Close     float64   `json:"close" gorm:"type:decimal(18,8);not null"`
	Volume    float64   `json:"volume" gorm:"type:decimal(18,8);not null"`
	Timeframe string    `json:"timeframe" gorm:"type:varchar(10);not null"`
}
