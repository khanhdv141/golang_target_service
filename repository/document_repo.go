package repository

import (
	"CMS/model"
	"context"
	"gorm.io/gorm"
)

type DocumentRepository interface {
	Save(context.Context, *model.Document) error
	Delete(context.Context, int64) error
	FindById(context.Context, int64) (*model.Document, error)
	FindAll(context.Context, int, int) ([]*model.Document, error)
}

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{db: db}
}

func (d *documentRepository) Save(ctx context.Context, document *model.Document) error {
	return d.db.WithContext(ctx).Save(document).Error
}

func (d *documentRepository) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Unscoped().Delete(&model.Document{}, id).Error
}

func (d *documentRepository) FindById(ctx context.Context, id int64) (*model.Document, error) {
	var document model.Document
	err := d.db.WithContext(ctx).Unscoped().Model(&model.Document{}).
		First(&document, id).Error
	if err != nil {
		return nil, err
	}
	return &document, err
}

func (d *documentRepository) FindAll(ctx context.Context, size int, page int) ([]*model.Document, error) {
	var documents []*model.Document
	err := d.db.WithContext(ctx).
		Model(&model.Document{}).
		Limit(size).
		Offset(page * size).
		Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, err
}
