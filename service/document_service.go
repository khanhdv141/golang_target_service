package service

import (
	"CMS/dto"
	"CMS/model"
	"CMS/repository"
	"github.com/gin-gonic/gin"
)

type DocumentService interface {
	CreateDocument(*gin.Context, dto.CreateDocumentRequest) *dto.BaseResponse[*model.Document]
	UpdateDocument(*gin.Context, dto.UpdateDocumentRequest) *dto.BaseResponse[*model.Document]
	DeleteDocument(*gin.Context, int64) *dto.BaseResponse[any]
	GetDocument(*gin.Context, int64) *dto.BaseResponse[*model.Document]
}

type documentService struct {
	documentRepository repository.DocumentRepository
	fileRepository     repository.FileRepository
	fileService        FileService
}

func NewDocumentService(
	documentRepository repository.DocumentRepository,
	fileRepository repository.FileRepository,
	fileService FileService) DocumentService {
	return &documentService{
		documentRepository: documentRepository,
		fileRepository:     fileRepository,
		fileService:        fileService,
	}
}

func (s *documentService) CreateDocument(ctx *gin.Context, req dto.CreateDocumentRequest) *dto.BaseResponse[*model.Document] {
	appFile := s.fileService.CreateFile(ctx, req.File)
	if appFile.Code != 200 {
		return MakeBadRequestResponse[*model.Document](appFile.Message)
	}

	document := model.Document{
		Name:            req.Name,
		Code:            req.Code,
		Type:            req.Type,
		IssuanceDate:    req.IssuanceDate,
		PublicationDate: req.PublicationDate,
		ExpirationDate:  req.ExpirationDate,
		EffectiveDate:   req.EffectiveDate,
		SourceFileId:    appFile.Data.ID,
		PreviewFileId:   appFile.Data.ID,
		EditableFileId:  appFile.Data.ID,
		// TODO: Assign created by
	}
	err := s.documentRepository.Save(ctx, &document)
	if err != nil {
		return MakeBadRequestResponse[*model.Document]("Cannot save document")
	}

	return MakeSuccessResponse[*model.Document](&document)
}

func (s *documentService) UpdateDocument(ctx *gin.Context, req dto.UpdateDocumentRequest) *dto.BaseResponse[*model.Document] {
	document, err := s.documentRepository.FindById(ctx, req.ID)
	if err != nil {
		return MakeBadRequestResponse[*model.Document]("Document with given ID does not exist")
	}
	if req.Name != "" {
		document.Name = req.Name
	}
	if req.File != nil {
		appFile := s.fileService.CreateFile(ctx, req.File)
		if appFile.Code != 200 {
			return MakeBadRequestResponse[*model.Document](appFile.Message)
		}
		document.PreviewFileId = appFile.Data.ID
		document.EditableFileId = appFile.Data.ID
		document.SourceFileId = appFile.Data.ID
	}
	if len(req.ParsedMetadata) != 0 {
		document.Metadata = req.ParsedMetadata
	}
	err = s.documentRepository.Save(ctx, document)
	if err != nil {
		return MakeBadRequestResponse[*model.Document]("Cannot save document")
	}
	return MakeSuccessResponse[*model.Document](document)
}

func (s *documentService) DeleteDocument(ctx *gin.Context, id int64) *dto.BaseResponse[any] {
	_, err := s.documentRepository.FindById(ctx, id)
	if err != nil {
		return MakeBadRequestResponse[any]("Document with given ID does not exist")
	}
	err = s.documentRepository.Delete(ctx, id)
	if err != nil {
		return MakeBadRequestResponse[any]("Cannot delete document")
	}
	return MakeSuccessResponseWithMessage[any](nil, "Deleted 1 document.")
}

func (s *documentService) GetDocument(ctx *gin.Context, id int64) *dto.BaseResponse[*model.Document] {
	document, err := s.documentRepository.FindById(ctx, id)
	if err != nil {
		return MakeBadRequestResponse[*model.Document]("Document with given ID does not exist")
	}
	return MakeSuccessResponse[*model.Document](document)
}

//func (s *documentService) ListDocuments(ctx *gin.Context) *dto.BaseResponse[*model.Document] {
//	var queryModel dto.ListDocumentRequest
//	err := ctx.ShouldBindQuery(&queryModel)
//	if err != nil {
//		return MakeBadRequestResponse[*model.Document]("Invalid query params")
//	}
//	if queryModel.Size == 0 {
//		queryModel.Size = 20
//	}
//	var filterConditions map[string]interface{}
//	if ()
//}
