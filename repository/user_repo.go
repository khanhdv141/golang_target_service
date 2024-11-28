package repository

import (
	"CMS/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(*gin.Context, *model.User) error
	Update(*gin.Context, *model.User) error
	DeleteById(*gin.Context, uint) error
	FindById(*gin.Context, uint) (*model.User, error)
	FindByUsername(*gin.Context, string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Save(ctx *gin.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) Update(ctx *gin.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) DeleteById(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (r *userRepo) FindById(ctx *gin.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepo) FindByUsername(ctx *gin.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
