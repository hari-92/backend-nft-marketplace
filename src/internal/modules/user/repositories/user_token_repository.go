package user_repositories

import (
	userFilters "gitlab.com/hari-92/nft-market-server/internal/modules/user/filters"
	userModels "gitlab.com/hari-92/nft-market-server/internal/modules/user/models"
	userScopes "gitlab.com/hari-92/nft-market-server/internal/modules/user/repositories/scopes"
	"gorm.io/gorm"
)

type IUserTokenRepository interface {
	FindOne(filter *userFilters.UserTokenFilter) (*userModels.UserToken, error)
	FindMany(filter *userFilters.UserTokenFilter) ([]*userModels.UserToken, error)
	Create(userToken *userModels.UserToken) (*userModels.UserToken, error)
	Update(filter *userFilters.UserTokenFilter, update interface{}) (*userModels.UserToken, error)
	Save(model *userModels.UserToken) (*userModels.UserToken, error)
}

func NewUserTokenRepository(
	db *gorm.DB,
) IUserTokenRepository {
	return &UserTokenRepository{
		db: db,
	}
}

type UserTokenRepository struct {
	db *gorm.DB
}

func (r *UserTokenRepository) FindOne(filter *userFilters.UserTokenFilter) (*userModels.UserToken, error) {
	var userToken userModels.UserToken
	err := r.db.Scopes(userScopes.UserTokenScope(filter)).First(&userToken).Error
	if err != nil {
		return nil, err
	}
	return &userToken, nil
}

func (r *UserTokenRepository) FindMany(filter *userFilters.UserTokenFilter) ([]*userModels.UserToken, error) {
	var userTokens []*userModels.UserToken
	err := r.db.Scopes(userScopes.UserTokenScope(filter)).Find(&userTokens).Error
	if err != nil {
		return nil, err
	}
	return userTokens, nil
}

func (r *UserTokenRepository) Create(userToken *userModels.UserToken) (*userModels.UserToken, error) {
	err := r.db.Create(&userToken).Error
	if err != nil {
		return nil, err
	}
	return userToken, nil
}

func (r *UserTokenRepository) Update(filter *userFilters.UserTokenFilter, update interface{}) (*userModels.UserToken, error) {
	var model *userModels.UserToken
	err := r.db.Scopes(userScopes.UserTokenScope(filter)).Updates(update).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *UserTokenRepository) Save(model *userModels.UserToken) (*userModels.UserToken, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
