package repository

import (
	"CMS/model"
	"context"
	"gorm.io/gorm"
)

type FileRepository interface {
	Save(ctx context.Context, file *model.File) error
	FindById(ctx context.Context, id string) (*model.File, error)
	FindByName(context.Context, string) (*model.File, error)
	DeleteById(ctx context.Context, id string) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) Save(ctx context.Context, file *model.File) error {
	return r.db.WithContext(ctx).Save(file).Error
}

func (r *fileRepository) FindById(ctx context.Context, id string) (*model.File, error) {
	var file model.File
	err := r.db.WithContext(ctx).First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) FindByName(ctx context.Context, name string) (*model.File, error) {
	var file model.File
	err := r.db.WithContext(ctx).First(&file, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) DeleteById(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.File{}, id).Error
}
