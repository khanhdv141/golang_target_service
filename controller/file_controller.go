package controller

import "CMS/service"

type FileController interface {
}

type fileController struct {
	fileService service.FileService
}

//func (c *fileController) CreateFile()
