package dto

import (
	"mime/multipart"
	"time"
)

type CreateDocumentRequest struct {
	Name            string                `form:"name" validate:"required"`
	File            *multipart.FileHeader `form:"file" validate:"required"`
	Code            string                `form:"code" validate:"required"`
	Type            string                `form:"type" validate:"required"`
	IssuanceDate    time.Time             `form:"issuance_date" validate:"required"`
	PublicationDate time.Time             `form:"publication_date"`
	ExpirationDate  time.Time             `form:"expiration_date"`
	EffectiveDate   time.Time             `form:"column:effective_date"`
	Metadata        string                `form:"metadata"`
	ParsedMetadata  map[string]interface{}
}

type UpdateDocumentRequest struct {
	ID              int64                 `form:"id" validate:"required"`
	Name            string                `form:"name" validate:"required"`
	File            *multipart.FileHeader `form:"file" validate:"required"`
	Code            string                `form:"code" validate:"required"`
	Type            string                `form:"type" validate:"required"`
	IssuanceDate    time.Time             `form:"issuance_date" validate:"required"`
	PublicationDate time.Time             `form:"publication_date"`
	ExpirationDate  time.Time             `form:"expiration_date"`
	EffectiveDate   time.Time             `form:"column:effective_date"`
	Metadata        string                `form:"metadata"`
	ParsedMetadata  map[string]interface{}
}

type ListDocumentRequest struct {
	Page           int    `form:"page"`
	Size           int    `form:"size"`
	Name           string `form:"name"`
	CreateTimeFrom string `form:"create_time_from"`
	CreateTimeTo   string `form:"create_time_to"`
}
