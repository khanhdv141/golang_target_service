package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID           string    `gorm:"primaryKey,column:id"`
	Name         string    `gorm:"column:name;uniqueIndex;size:255"`
	OriginalName string    `gorm:"column:original_name"`
	MimeType     string    `gorm:"column:mimetype"`
	Size         int64     `gorm:"column:size"`
	Path         string    `gorm:"column:path;uniqueIndex;size:300"`
	Extension    string    `gorm:"column:extension"`
	Type         string    `gorm:"column:type"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (file *File) BeforeCreate(*gorm.DB) (err error) {
	// UUID version 4
	if file.ID == "" {
		file.ID = uuid.NewString()
	}
	return
}
