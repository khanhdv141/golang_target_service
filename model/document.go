package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Document struct {
	ID              uint      `gorm:"primaryKey,column:id" json:"id"`
	CreatedAt       time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"updated_at" json:"updatedAt"`
	Name            string    `gorm:"column:name" json:"name"`
	Code            string    `gorm:"column:code" json:"code"`
	Type            string    `gorm:"column:type" json:"type"`
	IssuanceDate    time.Time `gorm:"column:issuance_date" json:"issuanceDate"`
	PublicationDate time.Time `gorm:"column:publication_date" json:"publicationDate"`
	ExpirationDate  time.Time `gorm:"column:expiration_date" json:"expirationDate"`
	EffectiveDate   time.Time `gorm:"column:effective_date" json:"effectiveDate"`
	SourceFileId    string    `gorm:"column:source_file_id" json:"sourceFileId"`
	PreviewFileId   string    `gorm:"column:preview_file_id" json:"previewFileId"`
	EditableFileId  string    `gorm:"column:editable_file_id" json:"editableFileId"`
	Metadata        Metadata  `gorm:"column:metadata;type:json" json:"metadata"`
	CreatedBy       string    `gorm:"column:created_by" json:"createdBy"`
}

type Metadata map[string]interface{}

func (m Metadata) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *Metadata) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}
