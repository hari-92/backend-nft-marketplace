package token_repositories

import (
	tokenFilters "gitlab.com/hari-92/nft-market-server/internal/modules/token/filters"
	tokenModels "gitlab.com/hari-92/nft-market-server/internal/modules/token/models"
	tokenScopes "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories/scopes"
	"gorm.io/gorm"
)

type TokenRepository interface {
	FindOne(filter *tokenFilters.TokenFilter) (*tokenModels.Token, error)
	FindMany(filter *tokenFilters.TokenFilter) ([]*tokenModels.Token, error)
	Create(model *tokenModels.Token) (*tokenModels.Token, error)
	Update(filter *tokenFilters.TokenFilter, update interface{}) (*tokenModels.Token, error)
	Save(model *tokenModels.Token) (*tokenModels.Token, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) FindOne(filter *tokenFilters.TokenFilter) (*tokenModels.Token, error) {
	var token tokenModels.Token
	err := r.db.Scopes(tokenScopes.TokenScope(filter)).First(&token).Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *tokenRepository) FindMany(filter *tokenFilters.TokenFilter) ([]*tokenModels.Token, error) {
	var tokens []*tokenModels.Token
	err := r.db.Scopes(tokenScopes.TokenScope(filter)).Find(&tokens).Error
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *tokenRepository) Create(model *tokenModels.Token) (*tokenModels.Token, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *tokenRepository) Update(filter *tokenFilters.TokenFilter, update interface{}) (*tokenModels.Token, error) {
	var model *tokenModels.Token
	err := r.db.Scopes(tokenScopes.TokenScope(filter)).Updates(update).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *tokenRepository) Save(model *tokenModels.Token) (*tokenModels.Token, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
