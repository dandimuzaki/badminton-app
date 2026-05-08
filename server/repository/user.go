package repository

import (
	"context"

	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, id uint) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{
		DB:  db,
		Log: log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	db := infra.GetDB(ctx, r.DB)
	var user model.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByID(ctx context.Context, id uint) (*model.User, error) {
	db := infra.GetDB(ctx, r.DB)
	var user model.User
	if err := db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(user).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}
