package service

import (
	"CMS/config"
	"CMS/dto"
	"CMS/model"
	"CMS/repository"
	"CMS/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileService interface {
	CreateFile(*gin.Context, *multipart.FileHeader) *dto.BaseResponse[*model.File]
}

type fileService struct {
	fileRepository repository.FileRepository
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileService{
		fileRepository: fileRepository,
	}
}

func (s *fileService) CreateFile(ctx *gin.Context, file *multipart.FileHeader) *dto.BaseResponse[*model.File] {
	tempDir := filepath.Join(os.TempDir(), uuid.New().String())
	_ = os.Mkdir(tempDir, os.ModePerm)
	filePath := filepath.Join(tempDir, file.Filename)

	err := ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		return MakeBadRequestResponse[*model.File]("Cannot read given file")
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return MakeBadRequestResponse[*model.File]("Cannot read given file")
	}

	stat, err := os.Stat(filePath)
	if err != nil {
		return MakeBadRequestResponse[*model.File]("Cannot read given file")
	}

	var fileModel model.File

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		fallthrough
	case "image/png":
		fallthrough
	case "image/gif":
		fallthrough
	case "image/svg+xml":
		fallthrough
	case "image/tiff":
		fallthrough
	case "image/webp":
		fileModel.Type = "image"
		break
	case "application/msword":
		fallthrough
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		fallthrough
	case "application/vnd.oasis.opendocument.text":
		fileModel.Type = "docx"
	case "application/pdf":
		fileModel.Type = "pdf"
	}

	fileModel.ID = uuid.New().String()
	fileModel.Size = stat.Size()
	fileModel.Extension = strings.Replace(filepath.Ext(filePath), ".", "", 1)
	fileModel.OriginalName = file.Filename
	fileModel.MimeType = mimeType
	fileModel.Name = fileModel.ID + "." + fileModel.Extension

	savePath := filepath.Join(
		config.ApplicationConfig.StorageDirectory,
		time.Now().Format("02-01-2006"))
	_ = os.MkdirAll(savePath, os.ModePerm)
	fileModel.Path = filepath.Join(savePath, fileModel.Name)

	err = util.CopyFile(filePath, fileModel.Path)
	if err != nil {
		fmt.Println(err)
		return MakeBadRequestResponse[*model.File]("Cannot save given file")
	}

	err = s.fileRepository.Save(ctx, &fileModel)
	if err != nil {
		fmt.Println(err)
		return MakeBadRequestResponse[*model.File]("Cannot save given file")
	}

	return MakeSuccessResponse[*model.File](&fileModel)
}
