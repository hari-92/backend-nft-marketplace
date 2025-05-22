package user_models

import (
	"time"

	"gitlab.com/hari-92/nft-market-server/internal/core/base"
)

type User struct {
	base.Model
	Username     string    `json:"username" gorm:"type:varchar(50);unique"`
	Email        string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	PasswordHash string    `json:"password_hash" gorm:"type:varchar(255)"`
	GoogleID     string    `json:"google_id" gorm:"type:varchar(100);unique"`
	PhoneNumber  string    `json:"phone_number" gorm:"type:varchar(20)"`
	Lang         string    `json:"lang" gorm:"type:varchar(2);default:en"`
	IsSetLang    bool      `json:"is_set_lang" gorm:"type:boolean;default:false"`
	LastLoginAt  time.Time `json:"last_login_at" gorm:"type:datetime"`
}
