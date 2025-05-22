package pnl_repositories

import (
	pnlFilters "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/filters"
	pnlModels "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/models"
	pnlScopes "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/repositories/scopes"
	"gorm.io/gorm"
)

type PnlRepository interface {
	FindOne(filter *pnlFilters.PnlFilter) (*pnlModels.PNL, error)
	FindMany(filter *pnlFilters.PnlFilter) ([]*pnlModels.PNL, error)
	Create(model *pnlModels.PNL) error
	Update(filter *pnlFilters.PnlFilter, model *pnlModels.PNL) error
}

func NewPnlRepository(
	db *gorm.DB,
) PnlRepository {
	return &pnlRepository{db: db}
}

type pnlRepository struct {
	db *gorm.DB
}

func (r *pnlRepository) FindOne(filter *pnlFilters.PnlFilter) (*pnlModels.PNL, error) {
	var model *pnlModels.PNL
	err := r.db.Model(&pnlModels.PNL{}).Scopes(pnlScopes.PnlScope(filter)).First(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *pnlRepository) FindMany(filter *pnlFilters.PnlFilter) ([]*pnlModels.PNL, error) {
	var models []*pnlModels.PNL
	err := r.db.Model(&pnlModels.PNL{}).Scopes(pnlScopes.PnlScope(filter)).Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *pnlRepository) Create(model *pnlModels.PNL) error {
	err := r.db.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *pnlRepository) Update(filter *pnlFilters.PnlFilter, model *pnlModels.PNL) error {
	err := r.db.Model(&pnlModels.PNL{}).Scopes(pnlScopes.PnlScope(filter)).Updates(model).Error
	if err != nil {
		return err
	}

	return nil
}
