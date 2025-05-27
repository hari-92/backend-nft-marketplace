package user_repositories

import (
	userFilters "gitlab.com/hari-92/nft-market-server/internal/modules/user/filters"
	userModels "gitlab.com/hari-92/nft-market-server/internal/modules/user/models"
	userScopes "gitlab.com/hari-92/nft-market-server/internal/modules/user/repositories/scopes"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetOne(filter *userFilters.UserFilter) (*userModels.User, error)
	GetMany(filter *userFilters.UserFilter) ([]*userModels.User, error)
	Create(model *userModels.User) (*userModels.User, error)
	Update(filter *userFilters.UserFilter, model *userModels.User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetOne(filter *userFilters.UserFilter) (*userModels.User, error) {
	var user userModels.User
	err := r.db.Scopes(userScopes.UserScope(filter)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetMany(filter *userFilters.UserFilter) ([]*userModels.User, error) {
	var users []*userModels.User
	err := r.db.Scopes(userScopes.UserScope(filter)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(model *userModels.User) (*userModels.User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *userRepository) Update(filter *userFilters.UserFilter, model *userModels.User) error {
	err := r.db.Scopes(userScopes.UserScope(filter)).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
